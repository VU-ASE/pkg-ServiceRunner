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
	Dependencies []ResolvedDependency // the dependencies of the service
}

// Utiliy function to get the address of a dependency
func (service ResolvedService) GetDependencyAddress(serviceName string, outputName string) (string, error) {
	for _, dependency := range service.Dependencies {
		if strings.ToLower(serviceName) == strings.ToLower(dependency.ServiceName) && strings.ToLower(outputName) == strings.ToLower(dependency.OutputName) {
			return dependency.Address, nil
		}
	}
	return "", fmt.Errorf("Dependency %s.%s not found", serviceName, outputName)
}
