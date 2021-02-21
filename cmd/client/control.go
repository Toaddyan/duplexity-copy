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
			ID: config.ClientID,
		},
	}
	controlMessageBytes, err := messages.MessageToControlMessageBytes(discoveryRequest)
	if err != nil {
		log.Panicf("%v", err)
	}
	sendChannel <- controlMessageBytes
}

// Need to establish user plane websocket
func processDiscoveryResponse() string {
	for {
		controlCommand := <-readChannel
		if controlCommand.MessageType == "DiscoveryResponse" {
			// message := messages.DiscoveryResponse{}
			message, err := messages.ControlMessageToMessage(controlCommand)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println(message)
			return ""
			// TODO Have to use this unmarshal method rather than reflect
			// err := json.Unmarshal([]byte(controlCommand.Message), &message)
			// if err != nil {
			// 	log.Panicf("%+v\n", err)
			// }
			// return message.DataPlaneURI
		}
	}
}
