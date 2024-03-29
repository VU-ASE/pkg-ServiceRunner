package servicerunner

import (
	"fmt"
	"strings"

	pb_systemmanager_messages "github.com/VU-ASE/pkg-CommunicationDefinitions/v2/packages/go/systemmanager"
	"github.com/rs/zerolog/log"
)

//
// This file contains the information that is passed back to the service that is run. The service can then use this information to connect to the dependencies.
//

// A resolved dependency, which will be given back to the service
type ResolvedDependency struct {
	ServiceName string // the name of the service that this service depends on
	OutputName  string // the name of the output that this service needs from the dependency
	Address     string // the address of the output that this service needs from the dependency
}

type ResolvedService struct {
	Name         string               // the name of the service
	Pid          int                  // the pid of the service
	Dependencies []ResolvedDependency // the dependencies of the service, already resolved
	Outputs      []output             // the outputs of this service
}

// For now, we only replace * with localhost for zmq, but more modifications can be added later
func rewriteDependencyAddress(addr string) string {
	depAddr := strings.ReplaceAll(addr, "*", "localhost")
	if depAddr != addr {
		log.Debug().Str("old", addr).Str("new", depAddr).Msg("Rewrote dependency address for own consumption")
	}
	return depAddr
}

// Utiliy function to get the address of a dependency
func (service ResolvedService) GetDependencyAddress(serviceName string, outputName string) (string, error) {
	for _, dependency := range service.Dependencies {
		if strings.EqualFold(serviceName, dependency.ServiceName) && strings.EqualFold(outputName, dependency.OutputName) {
			return rewriteDependencyAddress(dependency.Address), nil
		}
	}
	return "", fmt.Errorf("Dependency '%s.%s' not found. Are you sure it is exposed by %s?", serviceName, outputName, serviceName)
}

// For now, we only replace localhost with * for zmq, but more modifications can be added later
func rewriteOutputAddress(addr string) string {
	repAddr := strings.ReplaceAll(addr, "localhost", "*")
	if repAddr != addr {
		log.Debug().Str("old", addr).Str("new", repAddr).Msg("Rewrote output address for own consumption")
	}
	return repAddr
}

// Utility function to get the address of your own output
func (service ResolvedService) GetOutputAddress(outputName string) (string, error) {
	for _, output := range service.Outputs {
		if strings.EqualFold(outputName, output.Name) {
			return rewriteOutputAddress(output.Address), nil
		}
	}
	return "", fmt.Errorf("Output '%s' not found. Was it defined in service.yaml?", outputName)
}

// Information about the system manager. This struct has useful methods implemented to repeat the same operations on the dependencies.
type SystemManagerInfo struct {
	RepReqAddress    string // the req/rep address of the system manager
	BroadcastAddress string // the public broadcast address of the system manager
}

// Utility function to get a list of all services running on the system manager
func (sysman SystemManagerInfo) GetAllServices() (*pb_systemmanager_messages.ServiceList, error) {
	return getServiceList(sysman.RepReqAddress)
}

// Utility function to easily read values from the tuning state
func GetTuningInt(key string, tuningState *pb_systemmanager_messages.TuningState) (int, error) {
	if tuningState == nil {
		return 0, fmt.Errorf("Tuning state is nil")
	}

	// Iterate over all the tuning state values
	for _, tuningValue := range tuningState.DynamicParameters {
		val := tuningValue.GetInt()
		if val != nil && val.Key == key {
			return int(val.Value), nil
		}
	}

	return 0, fmt.Errorf("Key '%s' not found in tuning state", key)
}

// Utility function to easily read values from the tuning state
func GetTuningString(key string, tuningState *pb_systemmanager_messages.TuningState) (string, error) {
	if tuningState == nil {
		return "", fmt.Errorf("Tuning state is nil")
	}

	// Iterate over all the tuning state values
	for _, tuningValue := range tuningState.DynamicParameters {
		val := tuningValue.GetString_()
		if val != nil && val.Key == key {
			return val.Value, nil
		}
	}

	return "", fmt.Errorf("Key '%s' not found in tuning state", key)
}

// Utility function to easily read values from the tuning state
func GetTuningFloat(key string, tuningState *pb_systemmanager_messages.TuningState) (float32, error) {
	if tuningState == nil {
		return 0, fmt.Errorf("Tuning state is nil")
	}

	// Iterate over all the tuning state values
	for _, tuningValue := range tuningState.DynamicParameters {
		val := tuningValue.GetFloat()
		if val != nil && val.Key == key {
			return val.Value, nil
		}
	}

	return 0, fmt.Errorf("Key '%s' not found in tuning state", key)
}
