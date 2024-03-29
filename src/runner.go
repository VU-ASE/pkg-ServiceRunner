package servicerunner

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	build_debug "runtime/debug"

	pb_systemmanager_messages "github.com/VU-ASE/pkg-CommunicationDefinitions/v2/packages/go/systemmanager"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const SERVER_ADDR = "tcp://localhost:1337"

// The function that is called when a new tuning state is recevied
type TuningStateCallbackFunction func(tuningState *pb_systemmanager_messages.TuningState)

// The main function to run
type MainFunction func(serviceInformation ResolvedService, sysmanInformation SystemManagerInfo, initialTuningState *pb_systemmanager_messages.TuningState) error

// The system manager exposes two endpoints: a pub/sub endpoint for broadcasting service registration and a req/rep endpoint for registering services and resolving dependencies
// this struct is used to store the addresses of these endpoints

// This address should be set in the environment variable ASE_SYSMAN_SERVER_ADDRESS (for req/rep communication)
func getSystemManagerRepReqAddress() (string, error) {
	serverAddr := os.Getenv("ASE_SYSMAN_SERVER_ADDRESS")
	if serverAddr == "" {
		log.Warn().Msg(fmt.Sprintf("Environment variable ASE_SYSMAN_SERVER_ADDRESS is not set, using default address: %s", SERVER_ADDR))
	}
	serverAddr = SERVER_ADDR
	return serverAddr, nil
}

// Configures log level and output
func setupLogging(debug bool, outputPath string, service serviceDefinition) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	// Set up custom caller prefix
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		path := strings.Split(file, "/")
		// only take the last three elements of the path
		filepath := strings.Join(path[len(path)-3:], "/")
		return fmt.Sprintf("[%s] %s:%d", service.Name, filepath, line)
	}
	outputWriter := zerolog.ConsoleWriter{Out: os.Stderr}
	log.Logger = log.Output(outputWriter).With().Caller().Logger()
	if outputPath != "" {
		file, err := os.OpenFile(
			outputPath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			panic(err)
		}
		log.Logger = zerolog.New(file).With().Timestamp().Caller().Logger()
		fmt.Printf("Logging to file %s\n", outputPath)
	}

	// Set log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug logs enabled")
	}

	log.Info().Msg("Logger was set up")

	if debug {
		log.Debug().Msg("Listing dependencies of this binary: ")
		buildInfo, ok := build_debug.ReadBuildInfo()
		if !ok {
			log.Warn().Msg("Failed to read build info")
		} else {
			for _, dep := range buildInfo.Deps {
				s := fmt.Sprintf("  dep: %s@%s", dep.Path, dep.Version)
				log.Debug().Msg(s)
			}
		}
	}
}

// Used to start the program with the correct arguments and logging, with service discovery registration and all dependencies resolved
func Run(main MainFunction, onTuningState TuningStateCallbackFunction, disableRegistration bool) {
	// Parse args
	debug := flag.Bool("debug", false, "show all logs (including debug)")
	output := flag.String("output", "", "path of the output file to log to")
	retries := flag.Int("retries", 0, "how often to retry the service if it fails")
	serviceYamlPath := flag.String("service-yaml", "service.yaml", "path to the service definition yaml file")
	noLiveTuning := flag.Bool("disable-live-tuning", false, "disable live tuning updates from the system manager")

	flag.Parse()

	// Parse the service definition
	service, err := parseServiceDefinitionFromYaml(*serviceYamlPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal().Err(err).Msg("Could not open service definition YAML. Use the -service-yaml flag to specify the path to the service definition YAML file")
		} else {
			log.Fatal().Err(err).Msg("Error parsing service definition YAML")
		}
	}

	// Set up logging
	setupLogging(*debug, *output, service)

	// Try registering the service with the system manager
	resolvedDependencies := make([]ResolvedDependency, 0)

	// The address on which to send requests to the system manager
	// will be filled in according to the environment variable
	sysmanInfo := SystemManagerInfo{
		RepReqAddress:    "",
		BroadcastAddress: "",
	}

	// Don't register the system manager with itself
	if disableRegistration {
		log.Info().Msg("Service registration skipped was disabled by the user.")
	} else {
		// Where can we reach the system manager?
		sysmanInfo.RepReqAddress, err = getSystemManagerRepReqAddress()
		if err != nil {
			log.Fatal().Err(err).Msg("Error getting system manager details")
		}

		// Register the service with the system manager
		resolvedDependencies, err = registerService(service, sysmanInfo.RepReqAddress)
		if err != nil {
			log.Fatal().Err(err).Msg("Error registering service")
		}
	}

	// Callback information passed back to the service
	serviceInformation := ResolvedService{
		Name:         service.Name,
		Pid:          os.Getpid(),
		Dependencies: resolvedDependencies,
		Outputs:      service.Outputs,
	}

	// Receive the initial tuning state
	initialTuning, err := convertOptionsToTuningState(service.Options)
	if err != nil {
		log.Fatal().Err(err).Msg("Error converting options to tuning state")
	}
	if disableRegistration {
		log.Info().Msg("Network tuning state fetch skipped. Was disabled by the user.")
	} else {
		newTuning, err := requestTuningState(sysmanInfo.RepReqAddress)
		if err != nil {
			log.Fatal().Err(err).Msg("Error requesting tuning state")
		}
		initialTuning = mergeTuningStates(initialTuning, newTuning)
	}

	// We keep refetching the tuning state until we have resolved (at minimum) all options without default values
	unresolvedOptions := getUnsetOptions(initialTuning, service.Options)
	for len(unresolvedOptions) > 0 {
		log.Info().Msgf("Cannot start service yet. Waiting for %d nresolved option(s) to be resolved throug dynamic tuning. Retrying in 4 seconds.", len(unresolvedOptions))
		for _, opt := range unresolvedOptions {
			log.Info().Msgf("- Unresolved option: %s (of type %s)", opt.Name, opt.Type)
		}
		time.Sleep(4 * time.Second)
		newTuning, err := requestTuningState(sysmanInfo.RepReqAddress)
		if err != nil {
			log.Fatal().Err(err).Msg("Error requesting tuning state")
		}
		log.Info().Msg("Received new tuning state")
		initialTuning = mergeTuningStates(initialTuning, newTuning)
		unresolvedOptions = getUnsetOptions(initialTuning, service.Options)
	}

	if !*noLiveTuning && !disableRegistration {
		// We should be able to find the system manager broadcast address from our resolved dependencies
		sysmanInfo.BroadcastAddress, err = serviceInformation.GetDependencyAddress("systemmanager", "broadcast")
		if err != nil {
			log.Fatal().Err(err).Msg("Error getting system manager broadcast address")
		}

		// Listen for tuning state updates, and callback when a new tuning state is received
		go func() {
			for {
				err = listenForTuningBroadcasts(onTuningState, initialTuning, sysmanInfo.BroadcastAddress, service.Options)
				if err != nil {
					log.Err(err).Msg("Error listening for tuning state broadcasts")
				}
			}
		}()
	}

	// Identifier object to use for coming requests
	identifier := pb_systemmanager_messages.ServiceIdentifier{
		Name: service.Name,
		Pid:  int32(os.Getpid()),
	}

	backoff := 1
	// Always at least one try
	*retries++

	if *retries > 0 {
		*retries--
		log.Info().Msg("Starting service")
		if !disableRegistration {
			go func() {
				_ = updateServiceStatus(
					sysmanInfo.RepReqAddress,
					&identifier,
					pb_systemmanager_messages.ServiceStatus_RUNNING)
			}()
		}
		err = main(serviceInformation, sysmanInfo, initialTuning)
		if !disableRegistration {
			go func() {
				_ = updateServiceStatus(
					sysmanInfo.RepReqAddress,
					&identifier,
					pb_systemmanager_messages.ServiceStatus_STOPPED)
			}()
		}
		backoff *= 2
		log.Warn().Err(err).Int("retries left", *retries).Int("backoff", backoff).Msg("Service quit unexpectedly.")
		if *retries > 0 {
			log.Info().Msg("Retrying service...")
			time.Sleep(time.Duration(backoff) * time.Second)
		}
	}

	if err != nil {
		log.Err(err).Msg("Service quit unexpectedly, no retries left. Exiting...")
		os.Exit(1)
	} else {
		log.Info().Msg("Service finished successfully")
	}
}
