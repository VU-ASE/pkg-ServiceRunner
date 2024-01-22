package servicerunner

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

//
// This file describes the structure of a service definition, as should be included in the service.yaml file.
// The service.yaml file is used to describe the service and its dependencies and needs to follow this structure.
//

type ServiceDefinition struct {
	Name         string       `yaml:"name"`
	Description  string       `yaml:"description"`
	Dependencies []Dependency `yaml:"dependencies"`
	Outputs      []Output     `yaml:"outputs"`
}

type Dependency struct {
	// name of the service that this service depends on
	ServiceName string `yaml:"service"`
	// the name of the output that this service needs from the dependency
	OutputName string `yaml:"output"`
}

// The output that this service will produce
type Output struct {
	// the name of the output that this service will produce
	Name string `yaml:"name"`
	// the address of the output that this service will produce
	Address string `yaml:"address"`
}

func validateServiceDefinition(serviceDefinition ServiceDefinition) error {
	if serviceDefinition.Name == "" {
		return fmt.Errorf("Service name is empty")
	} else if serviceDefinition.Description == "" {
		return fmt.Errorf("Service description is empty")
	}

	if len(serviceDefinition.Dependencies) > 0 {
		for i, dependency := range serviceDefinition.Dependencies {
			if dependency.ServiceName == "" {
				return fmt.Errorf("Dependency service name is empty")
			} else if dependency.OutputName == "" {
				return fmt.Errorf("Dependency output name is empty")
			}

			// Check if service name and output name together are unique
			for j, otherDependency := range serviceDefinition.Dependencies {
				if i != j && dependency.ServiceName == otherDependency.ServiceName && dependency.OutputName == otherDependency.OutputName {
					return fmt.Errorf("Duplicate dependency: %s %s", dependency.ServiceName, dependency.OutputName)
				}
			}
		}
	}

	if len(serviceDefinition.Outputs) > 0 {
		for i, output := range serviceDefinition.Outputs {
			if output.Name == "" {
				return fmt.Errorf("Output name is empty")
			} else if output.Address == "" {
				return fmt.Errorf("Output address is empty")
			}

			// Check if names and addresses (individually) are unique
			for j, otherOutput := range serviceDefinition.Outputs {
				if i != j && output.Name == otherOutput.Name {
					return fmt.Errorf("Duplicate output name: %s", output.Name)
				}
				if i != j && output.Address == otherOutput.Address {
					return fmt.Errorf("Duplicate output address: %s", output.Address)
				}
			}
		}
	}

	return nil
}

func parseServiceDefinition(yamlString string) (ServiceDefinition, error) {
	serviceDefinition := ServiceDefinition{}
	err := yaml.Unmarshal([]byte(yamlString), &serviceDefinition)
	if err != nil {
		return ServiceDefinition{}, err
	}

	validationError := validateServiceDefinition(serviceDefinition)
	return serviceDefinition, validationError
}

func parseServiceDefinitionFromYaml(path string) (ServiceDefinition, error) {
	// Read file
	yaml, err := os.ReadFile(path)
	if err != nil {
		return ServiceDefinition{}, err
	}
	return parseServiceDefinition(string(yaml))
}
