package main

import (
	"encoding/json"
	"log"

	messagespb "github.com/duplexityio/duplexity/pkg/messages/pb"
	"github.com/gorilla/websocket"
)

func pingHandler(clientID string) {
	for {
		<-pingChannel
		ControlConnection.WriteMessage(websocket.PingMessage, []byte(clientID))
	}
}
func writePump() {
	for {
		controlMessage := <-sendChannel
		controlMessageBytes, err := json.Marshal(controlMessage)
		if err != nil {
			log.Panicf("%v\n", err)
		}
		ControlConnection.WriteMessage(websocket.TextMessage, controlMessageBytes)
	}
}

// readPump listens for any incoming messages through client-server websocket
// It then redirects these messages into the appropriate channels
func readPump() {
	for {
		mt, controlMessageBytes, err := ControlConnection.ReadMessage()
		if err != nil {
			log.Panicf("%v\n", err)
		}
		if mt == websocket.PingMessage {
			pingChannel <- true
		}
		if mt == websocket.TextMessage {
			controlMessage := messagespb.ControlMessage{}
			err := json.Unmarshal(controlMessageBytes, &controlMessage)
			if err != nil {
				log.Panicf("%v\n", err)
			}
			readChannel <- controlMessage
		}
	}
}
