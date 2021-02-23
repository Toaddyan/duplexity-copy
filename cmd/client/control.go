package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/pkg/messages"
	messagespb "github.com/duplexityio/duplexity/pkg/messages/pb"
	"github.com/gorilla/websocket"
	"github.com/rancher/remotedialer"
)

func dialControlConnection(controlWebsocketURI, clientID string) {
	var err error
	ControlConnection, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/control", ControlWebsocketURI), nil)
	if err != nil {
		log.Fatalf("Could not dial control websocket: %v\n", err)
	}
	return
}
func sendDiscoveryRequest(clientID string) {
	controlMessage := messagespb.ControlMessage{
		ClientID: clientID,
		MessageType: &messagespb.ControlMessage_DiscoveryRequest{
			DiscoveryRequest: &messagespb.DiscoveryRequest{
				Client: &messagespb.Client{
					ClientID: clientID,
				},
			},
		},
	}
	sendChannel <- controlMessage
}
func processDiscoveryResponse(clientID string) (string, error) {
	for {
		controlMessage := <-readChannel
		// Double Check who the message is intended for
		if controlMessage.ClientID != clientID {
			log.Println("Message is not intended for user")
			return "", errors.New("Broken Connection")
		}
		// Interpret what kind of message was sent.
		switch typeObject := controlMessage.MessageType.(type) {
		case *messagespb.ControlMessage_DiscoveryResponse:
			dataPlaneURI := typeObject.DiscoveryResponse.GetDataPlaneURI()
			return dataPlaneURI, nil
		default:
			log.Panic("Unable to process DiscoveryResponse...")
			continue
		}
	}
}

func buildPipes(ctx context.Context, clientID, dataPlaneURI, resource string) {
	headers := http.Header{
		messages.ClientIDHeaderKey: []string{clientID},
		messages.ResourceHeaderKey: []string{resource},
	}
	remotedialer.ClientConnect(ctx, dataPlaneURI, headers, nil, authorizer, nil)
}

func authorizer(protocol, address string) bool {
	return true
}

func listen() {
	for {

		controlMessage := <-readChannel
		switch typeObject := controlMessage.MessageType.(type) {
		// Check for Disconnects from the server
		case *messagespb.ControlMessage_Disconnect:
			if typeObject.Disconnect.Successful == true {
				log.Printf("You have been disconnected by the server %v")
				disconnect <- true
				return
			}
		default:
			log.Panic("Unable to process Request...")
			continue
		}
	}
}
