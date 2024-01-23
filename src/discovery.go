package servicerunner

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	protobuf_msgs "github.com/VU-ASE/pkg-ServiceRunner/include"
	customerrors "github.com/VU-ASE/pkg-ServiceRunner/src/errors"
	zmq "github.com/pebbe/zmq4"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

func allDependenciesResolved(service serviceDefinition, resolvedDependencies []ResolvedDependency) bool {
	if len(service.Dependencies) == 0 {
		return true
	}

	for _, dependency := range service.Dependencies {
		found := false
		for _, resolvedDependency := range resolvedDependencies {
			if dependency.ServiceName == resolvedDependency.ServiceName && dependency.OutputName == resolvedDependency.OutputName {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	return true
}

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

func extractOutputFromServiceStatus(status *protobuf_msgs.ServiceStatus, dependency dependency) (ResolvedDependency, error) {
	service := status.GetService()
	if service == nil {
		return ResolvedDependency{}, fmt.Errorf("Received empty service status")
	}

	// check if the service exposes the output that we need
	endpoints := service.GetEndpoints()
	if endpoints == nil {
		return ResolvedDependency{}, fmt.Errorf("Received empty service endpoints")
	}

	for _, endpoint := range endpoints {
		if endpoint.GetName() == dependency.OutputName {
			return ResolvedDependency{
				ServiceName: dependency.ServiceName,
				OutputName:  dependency.OutputName,
				Address:     endpoint.GetAddress(),
			}, nil
		}
	}

	return ResolvedDependency{}, customerrors.OutputNotExposed
}

// Will contact the discovery service to get the addresses of each dependency and register this service with the service discovery service (the system manager)
func registerService(service serviceDefinition) ([]ResolvedDependency, error) {
	// read the address of the system manager from the environment
	sysmanDetails, err := getSystemManagerDetails()
	if err != nil {
		return nil, err
	}

	// create a zmq client socket to the system manager
	client, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		return nil, fmt.Errorf("Could not open ZMQ connection to system manager: %s", err)
	}
	defer client.Close()
	log.Debug().Str("service", service.Name).Str("address", sysmanDetails.serverAddress).Msg("Connecting to system manager")
	err = client.Connect(sysmanDetails.serverAddress)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to system manager: %s", err)
	}

	// convert our service definition to a protobuf message
	endpoints := []*protobuf_msgs.ServiceEndpoint{}
	for _, output := range service.Outputs {
		// convert our struct to the ServiceEndpoint protobuf message
		endpoints = append(endpoints, &protobuf_msgs.ServiceEndpoint{
			Name:    output.Name,
			Address: output.Address,
		})
	}
	// create a registration message
	regMsg := protobuf_msgs.SystemManagerMessage{
		Msg: &protobuf_msgs.SystemManagerMessage_Service{
			Service: &protobuf_msgs.Service{
				Identifier: &protobuf_msgs.ServiceIdentifier{
					Name: service.Name,
					Pid:  int32(os.Getpid()),
				},
				Endpoints: endpoints,
			},
		},
	}

	// convert the message to bytes
	msgBytes, err := proto.Marshal(&regMsg)
	if err != nil {
		log.Err(err).Msg("Error marshalling protobuf message")
		return nil, err
	}

	// send registration to the system manager
	log.Info().Str("service", service.Name).Msg("Registering service with system manager")
	_, err = client.SendBytes(msgBytes, 0)
	if err != nil {
		return nil, err
	}

	responseReceived := false
	go func() {
		for {
			// print a idle message every 5 seconds, until were done
			if responseReceived {
				return
			}
			log.Info().Str("service", service.Name).Msg("Waiting for response from system manager")
			time.Sleep(5 * time.Second)
		}
	}()

	// wait for a response from the system manager
	resBytes, err := client.RecvBytes(0)
	responseReceived = true

	// the response must be of type Service (see messages/servicediscovery.proto)
	// if not, we discard it: registration not successful
	log.Info().Str("service", service.Name).Msg("Received registration response from system manager")
	if err != nil {
		return nil, err
	}
	response := protobuf_msgs.SystemManagerMessage{}
	err = proto.Unmarshal(resBytes, &response)
	if err != nil {
		log.Err(err).Msg("Error unmarshalling protobuf message")
		return nil, err
	}
	responseService := response.GetService()
	if responseService == nil {
		return nil, fmt.Errorf("Received empty response from system manager")
	}
	// check if the name and pid of the response match our registration, if not someone else registered with the same name
	identifier := responseService.GetIdentifier()
	if identifier == nil {
		return nil, fmt.Errorf("Received empty response from system manager")
	}
	name := identifier.GetName()
	pid := identifier.GetPid()
	if name != service.Name {
		return nil, fmt.Errorf("System manager denied service registration, name mismatch (registered as %s, but received %s)", service.Name, name)
	}
	if pid != int32(os.Getpid()) {
		return nil, fmt.Errorf("System manager denied service registration, service %s was already registered by pid %d", service.Name, pid)
	}
	// check if the endpoints match our registration
	registeredEndpints := responseService.GetEndpoints()
	if endpoints != nil {
		for i, endpoint := range endpoints {
			registeredEndpoint := registeredEndpints[i]
			if registeredEndpoint == nil {
				return nil, fmt.Errorf("Endpoint %s was not registered", endpoint.Name)
			} else if registeredEndpoint.GetName() != endpoint.Name || registeredEndpoint.GetAddress() != endpoint.Address {
				return nil, fmt.Errorf("Endpoint %s was registered with different address (%s) than requested (%s)", endpoint.Name, registeredEndpoint.GetAddress(), endpoint.Address)
			}
		}
	}

	// registration was successfull!
	log.Info().Str("service", service.Name).Msg("Service registration successful")

	resolvedDependencies := make([]ResolvedDependency, 0)
	// Now attempt to resolve dependencies, if any
	if len(service.Dependencies) > 0 {
		log.Info().Str("service", service.Name).Int("dependencies to resolve", len(service.Dependencies)).Msg("Resolving dependencies")
		for !allDependenciesResolved(service, resolvedDependencies) {
			// create a list all unique *services* (not endpoints) that we depend on and that are not yet resolved
			uniqueServiceDependencies := make([]string, 0)
			for _, dependency := range service.Dependencies {
				if !slices.Contains(uniqueServiceDependencies, dependency.ServiceName) && !slices.ContainsFunc(resolvedDependencies, func(dep ResolvedDependency) bool {
					return strings.ToLower(dep.ServiceName) == strings.ToLower(dependency.ServiceName) && strings.ToLower(dep.OutputName) == strings.ToLower(dependency.OutputName)
				}) {
					uniqueServiceDependencies = append(uniqueServiceDependencies, dependency.ServiceName)
				}
			}

			// resolve each service (sequentially)
			for _, serviceName := range uniqueServiceDependencies {
				// create a request message
				reqMsg := protobuf_msgs.SystemManagerMessage{
					Msg: &protobuf_msgs.SystemManagerMessage_ServiceInformationRequest{
						ServiceInformationRequest: &protobuf_msgs.ServiceInformationRequest{
							Requested: &protobuf_msgs.ServiceIdentifier{
								Name: serviceName,
								Pid:  1, // does not matter
							},
						},
					},
				}
				// convert the message to bytes
				msgBytes, err := proto.Marshal(&reqMsg)
				if err != nil {
					log.Err(err).Msg("Error marshalling protobuf message")
					return nil, err
				}
				// send the request to the system manager
				_, err = client.SendBytes(msgBytes, 0)
				if err != nil {
					return nil, err
				}
				log.Info().Str("service", service.Name).Str("dependency", serviceName).Msg("Requesting dependency information from system manager")
				gotReply := false
				go func() {
					for {
						// print a idle message every 5 seconds, until were done
						if gotReply {
							return
						}
						log.Info().Str("service", service.Name).Str("dependency", serviceName).Msg("Waiting for dependency response from system manager")
						time.Sleep(5 * time.Second)
					}
				}()
				// wait for a response from the system manager
				resBytes, err := client.RecvBytes(0)
				gotReply = true
				if err != nil {
					return nil, err
				}
				// the response must be of type ServiceStatus (see messages/servicediscovery.proto)
				response := protobuf_msgs.SystemManagerMessage{}
				responseServiceStatus := response.GetServiceStatus()
				if responseServiceStatus == nil {
					return nil, fmt.Errorf("Received empty response from system manager")
				}
				err = proto.Unmarshal(resBytes, &response)
				if err != nil {
					log.Warn().Str("service", service.Name).Str("dependency", serviceName).Err(err).Msg("Could not parse response from system manager. Assuming that the service we are depending on is not yet registered. Will retry in 5 seconds")
					time.Sleep(5 * time.Second)
					continue
				} else if responseServiceStatus.Status.Enum() != protobuf_msgs.ServiceStatus_RUNNING.Enum() {
					log.Warn().Str("service", service.Name).Str("dependency", serviceName).Msg("Dependency is not running yet. Will retry in 5 seconds")
					time.Sleep(5 * time.Second)
					continue
				}
				// check if this service exposes the outputs that we need
				for _, dependency := range service.Dependencies {
					if strings.ToLower(dependency.ServiceName) == strings.ToLower(serviceName) {
						// does the service expose our dependency as an output?
						resolvedDependency, err := extractOutputFromServiceStatus(responseServiceStatus, dependency)
						if err != nil {
							log.Err(err).Msg("Error extracting dependency from service status")
							return nil, err
						}
						// Add the resolved dependency to the list of resolved dependencies, if it is not already in there
						if !slices.ContainsFunc(resolvedDependencies, func(dep ResolvedDependency) bool {
							return strings.ToLower(dep.ServiceName) == strings.ToLower(resolvedDependency.ServiceName) && strings.ToLower(dep.OutputName) == strings.ToLower(resolvedDependency.OutputName)
						}) {
							resolvedDependencies = append(resolvedDependencies, resolvedDependency)
						}
					}

				}

			}
		}
		return resolvedDependencies, nil
	} else {
		log.Info().Str("service", service.Name).Msg("No dependencies to resolve")
		return make([]ResolvedDependency, 0), nil
	}
}
