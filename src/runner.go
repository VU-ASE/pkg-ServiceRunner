package servicerunner

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// The main function to run
type MainFunction func() error

// Configures log level and output
func setupLogging(debug bool, outputPath string, service ServiceDefinition) {
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

// Used to start the program with the correct arguments
func Run(main MainFunction) {
	// Parse args
	debug := flag.Bool("debug", false, "show all logs (including debug)")
	output := flag.String("output", "", "path of the output file to log to")
	retries := flag.Int("retries", 0, "how often to retry the service if it fails")
	serviceYamlPath := flag.String("service-yaml", "service.yaml", "path to the service definition yaml file")
	flag.Parse()

	// Parse the service definition
	service, err := ParseServiceDefinitionFromYaml(*serviceYamlPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing service definition")
	}

	// Set up logging
	setupLogging(*debug, *output, service)
	log.Info().Str("service", service.Name).Msg("Starting service")

	backoff := 1
	log.Info().Msg("Starting service")
	err = main()
	if *retries > 0 {
		*retries--
		backoff *= 2
		log.Warn().Err(err).Int("retries left", *retries).Int("backoff", backoff).Msg("Service quit unexpectedly. Retrying... ")
		time.Sleep(time.Duration(backoff) * time.Second)
		log.Info().Msg("Starting service")
		err = main()
	}

	if err != nil {
		log.Err(err).Msg("Service quit unexpectedly, no retries left. Exiting...")
		os.Exit(1)
	} else {
		log.Info().Msg("Service finished successfully")
	}
}
