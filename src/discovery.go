package servicerunner

import (
	"fmt"
	"os"

	protobuf_msgs "github.com/VU-ASE/pkg-ServiceRunner/src/messages"
	zmq "github.com/pebbe/zmq4"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type ResolvedDependency struct {
	ServiceName string // the name of the service that this service depends on
	OutputName  string // the name of the output that this service needs from the dependency
	Address     string // the address of the output that this service needs from the dependency
}

func allDependenciesResolved(service ServiceDefinition, resolvedDependencies []ResolvedDependency) bool {
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

// Will contact the discovery service to get the addresses of each dependency and register this service with the service discovery service (the system manager)
func registerService(service ServiceDefinition) ([]ResolvedDependency, error) {
	// read the address of the system manager from the environment
	sysmanDetails, err := getSystemManagerDetails()
	if err != nil {
		return nil, err
	}

	// create a zmq client socket to the system manager
	client, err := zmq.NewSocket(zmq.REP)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	err = client.Connect(sysmanDetails.serverAddress)
	if err != nil {
		return nil, err
	}

	// convert our service definition to a protobuf message
	endpoints := []*protobuf_msgs.ServiceEndpoint{}
	for _, output := range service.Outputs {
		endpoints = append(endpoints, &protobuf_msgs.ServiceEndpoint{
			Name:    output.Name,
			Address: output.Address,
		})
	}
	regMsg := protobuf_msgs.SystemManagerMessage{
		Message: &protobuf_msgs.SystemManagerMessage_ServiceDiscovery{
			ServiceDiscovery: &protobuf_msgs.ServiceDiscovery{
				Content: &protobuf_msgs.ServiceDiscovery_Registration_{
					Registration: &protobuf_msgs.ServiceDiscovery_Registration{
						Service: &protobuf_msgs.Service{
							Identifier: &protobuf_msgs.ServiceIdentifier{
								Name: service.Name,
								Pid:  int32(os.Getpid()),
							},
							Endpoints: endpoints,
						},
					},
				},
			},
		},
	}
	msgBytes, err := proto.Marshal(&regMsg)
	if err != nil {
		log.Err(err).Msg("Error marshalling protobuf message")
		return nil, err
	}

	// register our service with the system manager
	log.Info().Str("service", service.Name).Msg("Registering service with system manager")
	_, err = client.SendBytes(msgBytes, 0)
	if err != nil {
		return nil, err
	}

	// wait for a response from the system manager, if this is no error, we can assume that the service was registered successfully
	resBytes, err := client.RecvBytes(0)
	log.Info().Str("service", service.Name).Msg("Received first response from system manager, waiting for broadcast")
	if err != nil {
		return nil, err
	}
	response := protobuf_msgs.SystemManagerMessage{}
	response.Kin
	err = proto.Unmarshal(resBytes, &response)
	if err != nil {
		return nil, err
	}
	if response.GetMessage() == nil {
		return nil, fmt.Errorf("Received empty response from system manager")
	}
	resError := response.Message.(*protobuf_msgs.SystemManagerMessage_ServiceDiscovery).ServiceDiscovery.Content.(*protobuf_msgs.ServiceDiscovery_Error_).Error
	if resError != nil {
		return nil, fmt.Errorf("System manager denied service registration")
	}
	log.Info().Str("service", service.Name).Msg("Service registration successful")

	// todo: resolve dependencies by checking

}
