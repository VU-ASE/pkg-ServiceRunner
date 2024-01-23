package servicerunner

import (
	"fmt"
	"time"

	protobuf_msgs "github.com/VU-ASE/pkg-ServiceRunner/include"
	zmq "github.com/pebbe/zmq4"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

// Listens on the broadcast channel for broadcasts from the system manager
func listenForTuningBroadcasts(onTuningState TuningStateCallbackFunction, sysmanDetails SystemManagerDetails) error {
	// Subscribe to the system manager
	broadcast, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return err
	}
	defer broadcast.Close()
	broadcast.Connect(sysmanDetails.publisherAddress)
	broadcast.SetSubscribe("")

	// main receiver loop
	for {
		// Receive the tuning state
		msg, err := broadcast.RecvBytes(0)
		if err != nil {
			return err
		}

		// NOTE: this is a best effor, we discard the message (and any errors) if we cannot parse it
		// Parse the tuning state
		parsedMsg := protobuf_msgs.SystemManagerMessage{}
		err = proto.Unmarshal(msg, &parsedMsg)
		if err != nil {
			continue
		}

		tuningState := parsedMsg.GetTuningState()
		if tuningState == nil {
			continue
		}

		// Send the tuning state to the callback function
		log.Debug().Msg("Received tuning state broadcast from system manager")
		onTuningState(tuningState)
	}
}

func requestTuningState(sysmanDetails SystemManagerDetails) (*protobuf_msgs.TuningState, error) {
	// create a zmq client socket to the system manager
	client, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		return nil, fmt.Errorf("Could not open ZMQ connection to system manager: %s", err)
	}
	defer client.Close()
	log.Debug().Str("address", sysmanDetails.serverAddress).Msg("Connecting to system manager to fetch tuning state")
	err = client.Connect(sysmanDetails.serverAddress)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to system manager: %s", err)
	}

	// create a request message
	reqMsg := protobuf_msgs.SystemManagerMessage{
		Msg: &protobuf_msgs.SystemManagerMessage_TuningStateRequest{
			TuningStateRequest: &protobuf_msgs.TuningStateRequest{},
		},
	}

	// convert the message to bytes
	msgBytes, err := proto.Marshal(&reqMsg)
	if err != nil {
		log.Err(err).Msg("Error marshalling protobuf message")
		return nil, err
	}

	// send request to the system manager
	log.Info().Msg("Requesting tuning state from system manager")
	_, err = client.SendBytes(msgBytes, 0)
	if err != nil {
		return nil, err
	}

	responseReceived := false
	go func() {
		count := 0
		for {
			// print a idle message every 5 seconds, until were done
			if responseReceived {
				return
			}
			if (count) > 5 {
				log.Warn().Msg("Still waiting for tuning state response from system manager. Are you sure the system manager is running ?")
			} else {
				log.Info().Msg("Waiting for tuning state response from system manager")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// wait for a response from the system manager
	resBytes, err := client.RecvBytes(0)
	responseReceived = true

	// the response must be of type TuningState (see include/tuningstate.proto)
	// if not, we discard it: registration not successful
	log.Info().Msg("Received tuning state response from system manager")
	if err != nil {
		return nil, err
	}
	response := protobuf_msgs.SystemManagerMessage{}
	err = proto.Unmarshal(resBytes, &response)
	if err != nil {
		log.Err(err).Msg("Error unmarshalling tuning state protobuf message")
		return nil, err
	}
	responseState := response.GetTuningState()
	if responseState == nil {
		return nil, fmt.Errorf("Received empty response from system manager")
	}

	return responseState, nil
}
