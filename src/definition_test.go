package servicerunner

import (
	"fmt"
	"testing"
)

// This YAML file should be parsed correctly
func TestValidYaml(t *testing.T) {
	yamlPath := "../tests/valid.yaml"
	serviceDefinition, err := parseServiceDefinitionFromYaml(yamlPath)
	if err != nil {
		t.Errorf("Error parsing yaml: %s", err)
	} else {
		fmt.Printf("Service definition: %+v", serviceDefinition)
	}
}

// This YAML file should error
func TestInvalidPath(t *testing.T) {
	yamlPath := "../tests/doesnotexist.yaml"
	_, err := parseServiceDefinitionFromYaml(yamlPath)
	if err == nil {
		t.Errorf("Parsed invalid path")
	}
}

func TestInvalidYamlWithValuesMissing(t *testing.T) {
	yamlPath := "../tests/invalid-missing.yaml"
	_, err := parseServiceDefinitionFromYaml(yamlPath)
	if err == nil {
		t.Errorf("Parsed invalid yaml")
	}
}

func TestInvalidYamlWithDuplicateNames(t *testing.T) {
	yamlPath := "../tests/invalid-duplicate-names.yaml"
	_, err := parseServiceDefinitionFromYaml(yamlPath)
	if err == nil {
		t.Errorf("Parsed invalid yaml")
	}
}

func TestInvalidYamlWithDuplicateAddresses(t *testing.T) {
	yamlPath := "../tests/invalid-duplicate-addresses.yaml"
	_, err := parseServiceDefinitionFromYaml(yamlPath)
	if err == nil {
		t.Errorf("Parsed invalid yaml")
	}
}

func TestBasicRegistration(t *testing.T) {
	yamlPath := "../tests/basicregistration.yaml"
	serviceDefinition, err := parseServiceDefinitionFromYaml(yamlPath)
	if err != nil {
		t.Error(err)
	}

	t.Setenv("ASE_SYSMAN_SERVER_ADDRESS", "tcp://localhost:1337")

	_, err = registerService(serviceDefinition)
	if err != nil {
		t.Error(err)
	}
}
