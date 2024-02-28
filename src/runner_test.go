package servicerunner

import (
	"testing"

	"github.com/rs/zerolog/log"
)

// Show logs
func TestLogging(t *testing.T) {
	setupLogging(false, "", serviceDefinition{
		Name: "MyTestService",
	})
	log.Info().Msg("This is a test log")
}
