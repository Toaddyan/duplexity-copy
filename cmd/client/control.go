package main

import (
	"fmt"
	"log"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/websocket"
)

func dialControlConnection(ControlWebsocketURI, clientID string) {
	var err error
	ControlConnection, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/control", ControlWebsocketURI), nil)
	if err != nil {
		log.Fatalf("Could not dial control websocket: %v\n", err)
	}
	return
}

// sendDiscoveryRequest sends a request to the server to set up a data plane for client
func sendDiscoveryRequest(clientID string) {
	discoveryRequest := messages.DiscoveryRequest{
		Client: &messages.Client{
			config.ClientID,
		},
	}
	controlMessageBytes, err := messages.MessageToControlMessageBytes(discoveryRequest)
	sendChannel <- controlMessageBytes
}

// Need to establish user plane websocket
func processDiscoveryResponse() string {
	for {
		controlCommand := <-readChannel
		if controlCommand.MessageType == "DiscoveryResponse" {
			// message := messages.DiscoveryResponse{}
			message, err := messages.ControlMessageToMessage(controlCommand.Message)
			return message.DataPlaneURI
		}
	}
}
