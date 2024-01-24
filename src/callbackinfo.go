package servicerunner

import (
	"fmt"
	"strings"
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

// Utiliy function to get the address of a dependency
func (service ResolvedService) GetDependencyAddress(serviceName string, outputName string) (string, error) {
	for _, dependency := range service.Dependencies {
		if strings.EqualFold(serviceName, dependency.ServiceName) && strings.EqualFold(outputName, dependency.OutputName) {
			return dependency.Address, nil
		}
	}
	return "", fmt.Errorf("Dependency '%s.%s' not found. Are you sure it is exposed by %s?", serviceName, outputName, serviceName)
}

// Utility function to get the address of your own output
func (service ResolvedService) GetOutputAddress(outputName string) (string, error) {
	for _, output := range service.Outputs {
		if strings.EqualFold(outputName, output.Name) {
			return output.Address, nil
		}
	}
	return "", fmt.Errorf("Output '%s' not found. Was it defined in service.yaml?", outputName)
}
