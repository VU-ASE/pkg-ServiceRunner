package servicerunner

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	protobuf_msgs "github.com/VU-ASE/pkg-ServiceRunner/include"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// The function that is called when a new tuning state is recevied
type TuningStateCallbackFunction func(tuningState *protobuf_msgs.TuningState)

// The main function to run
type MainFunction func(serviceInformation ResolvedService, initialTuningState *protobuf_msgs.TuningState) error

// The system manager exposes two endpoints: a pub/sub endpoint for broadcasting service registration and a req/rep endpoint for registering services and resolving dependencies
// this struct is used to store the addresses of these endpoints
type SystemManagerDetails struct {
	serverAddress    string // used for req/rep communication
	publisherAddress string // used for pub/sub communication
}

func getSystemManagerDetails() (SystemManagerDetails, error) {
	serverAddr := os.Getenv("ASE_SYSMAN_SERVER_ADDRESS")
	if serverAddr == "" {
		return SystemManagerDetails{}, fmt.Errorf("Cannot reach system manager: environment variable ASE_SYSMAN_SERVER_ADDRESS is not set, do not know how to reach system manager :(")
	}
	pubAddr := os.Getenv("ASE_SYSMAN_BROADCAST_ADDRESS")
	if serverAddr == "" {
		return SystemManagerDetails{}, fmt.Errorf("Cannot reach system manager: environment variable ASE_SYSMAN_BROADCAST_ADDRESS is not set, do not know how to reach system manager :(")
	}
	return SystemManagerDetails{
		serverAddress:    serverAddr,
		publisherAddress: pubAddr,
	}, nil
}

// Configures log level and output
func setupLogging(debug bool, outputPath string, service serviceDefinition) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	// Log to stderr or to file
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if outputPath != "" {
		file, err := os.OpenFile(
			outputPath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			panic(err)
		}
		log.Logger = zerolog.New(file).With().Timestamp().Logger()
		fmt.Printf("Logging to file %s\n", outputPath)
	}

	// Set log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug logs enabled")
	}

	log.Info().Msg("Logger was set up")
}

// Used to start the program with the correct arguments and logging, with service discovery registration and all dependencies resolved
func Run(main MainFunction, onTuningState TuningStateCallbackFunction) {
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
			log.Fatal().Err(err).Msg("Error parsing service definition")
		}
	}

	// Set up logging
	setupLogging(*debug, *output, service)

	// Try registering the service with the system manager
	resolvedDependencies := make([]ResolvedDependency, 0)

	// Fetch the endpoints of the system manager
	systemManagerDetails, err := getSystemManagerDetails()
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting system manager details")
	}

	// Don't register the system manager with itself
	if strings.ToLower(service.Name) == "systemmanager" {
		log.Info().Msg("Service registration skipped. This is the system manager.")
	} else {
		// Register the service with the system manager
		resolvedDependencies, err = registerService(service, systemManagerDetails)
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
	tuningState := &protobuf_msgs.TuningState{}
	if strings.ToLower(service.Name) == "systemmanager" {
		log.Info().Msg("Tuning state skipped. This is the system manager.")
	} else {
		tuningState, err = requestTuningState(systemManagerDetails)
		if err != nil {
			log.Fatal().Err(err).Msg("Error requesting tuning state")
		}
	}

	log.Info().Str("service", service.Name).Msg("Starting service")
	backoff := 1
	log.Info().Msg("Starting service")
	err = main(serviceInformation, tuningState)
	if *retries > 0 {
		*retries--
		backoff *= 2
		log.Warn().Err(err).Int("retries left", *retries).Int("backoff", backoff).Msg("Service quit unexpectedly. Retrying... ")
		time.Sleep(time.Duration(backoff) * time.Second)
		log.Info().Msg("Starting service")
		err = main(serviceInformation, tuningState)
	}

	if !*noLiveTuning && strings.ToLower(service.Name) == "systemmanager" {
		// Listen for tuning state updates, and callback when a new tuning state is received
		go listenForTuningBroadcasts(onTuningState, systemManagerDetails)
	}

	if err != nil {
		log.Err(err).Msg("Service quit unexpectedly, no retries left. Exiting...")
		os.Exit(1)
	} else {
		log.Info().Msg("Service finished successfully")
	}
}
