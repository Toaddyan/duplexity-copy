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

// TODO globals
func dialControlConnection(controlWebsocketURI, clientID string) {
	var err error
	ControlConnection, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/control", controlWebsocketURI), nil)
	if err != nil {
		log.Fatalf("Could not dial control websocket: %v\n", err)
	}

	return
}

func sendRequest(cmd string) error {
	// controlMessage, err := messagespb.WrapCommand(req, config.ClientID)
	// if err != nil {
	// 	return err
	// }
	controlMessage := &messagespb.ControlMessage{}
	switch cmd {
	case "disconnectRequest":

		controlMessage = &messagespb.ControlMessage{
			ClientID:    config.ClientID,
			MessageType: &messagespb.ControlMessage_DisconnectRequest{
				DisconnectRequest: &messagespb.DisconnectRequest{},
			},
		}

	case "discoveryRequest":
		log.Println("Wrapping DiscoveryRequest")
		controlMessage = &messagespb.ControlMessage{
			ClientID: config.ClientID,
			MessageType: &messagespb.ControlMessage_DiscoveryRequest{
				DiscoveryRequest: &messagespb.DiscoveryRequest{},
			},
		}
	case "pipesRequest":
		controlMessage = &messagespb.ControlMessage{
			ClientID:    config.ClientID,
			MessageType: &messagespb.ControlMessage_PipesRequest{
				PipesRequest: &messagespb.PipesRegisteredRequest{
					Resource: resource,
				},
			},
		}
	default:
		return errors.New("wrong command ")
	}

	sendChannel <- controlMessage
	return nil
}

func unlock(lock bool) error {
	if lock == true {
		lock = false
		return nil
	}
	return errors.New("Lock has been previously unlocked")
}

func buildPipes(ctx context.Context, dataPlaneURI, resource string) {
	headers := http.Header{
		messages.ClientIDHeaderKey: []string{config.ClientID},
		messages.ResourceHeaderKey: []string{resource},
	}
	remotedialer.ClientConnect(ctx, dataPlaneURI, headers, nil, authorizer, nil)
}

func authorizer(protocol, address string) bool {
	return true
}

func listen() (interface{}, error) {
	for {
		controlMessage := <-readChannel
		// Double Check who the message is intended for
		if controlMessage.ClientID != config.ClientID {
			log.Println("Message is not intended for user")
			return "", errors.New("Broken Connection")
		}
		switch typeObject := controlMessage.MessageType.(type) {
		// Server has responded that control hub has registered client
		case *messagespb.ControlMessage_DiscoveryResponse:
			dataPlaneURI := typeObject.DiscoveryResponse.GetDataPlaneURI()

			err := unlock(controlLock)
			if err != nil {
				return "", err
			}
			return dataPlaneURI, nil
		// Server has responded that backend pipes have been established
		case *messagespb.ControlMessage_PipesResponse:
			pipeStatus := typeObject.PipesResponse.GetSuccessful()

			err := unlock(pipeLock)
			if err != nil {
				return "", err
			}

			return pipeStatus, nil
		// Check for Disconnects from the server
		case *messagespb.ControlMessage_DisconnectResponse:
			log.Printf("You have been disconnected by the server")
			disconnect <- true
		default:
			log.Panic("Unable to process Request...")
			continue
		}
	}
}
