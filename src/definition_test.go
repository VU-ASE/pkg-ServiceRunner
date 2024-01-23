package servicerunner

import (
	"fmt"
	"testing"
	"time"
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
	sysmanDetails := SystemManagerDetails{
		serverAddress:    "tcp://localhost:1337",
		publisherAddress: "tcp://localhost:1338",
	}

	done := make(chan bool)
	// first, boot up service A, which depends on B
	go func() {
		yamlPath := "../tests/dependencytesting/A.yaml"
		serviceDefinition, err := parseServiceDefinitionFromYaml(yamlPath)
		if err != nil {
			t.Error(err)
		}

		_, err = registerService(serviceDefinition, sysmanDetails)
		if err != nil {
			t.Error(err)
		}
		done <- true
	}()

	// wait 10 seconds, then boot up service B
	// this should also complete the registration of service A
	go func() {
		time.Sleep(10 * time.Second)

		yamlPath := "../tests/dependencytesting/B.yaml"
		serviceDefinition, err := parseServiceDefinitionFromYaml(yamlPath)
		if err != nil {
			t.Error(err)
		}

		_, err = registerService(serviceDefinition, sysmanDetails)
		if err != nil {
			t.Error(err)
		}
		done <- true
	}()

	// wait for both services to register
	<-done
	<-done
}
