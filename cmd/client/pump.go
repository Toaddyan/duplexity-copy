package main

import (
	"log"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/websocket"
)

func pingHandler(clientID string) {
	for {
		<-pingChannel
		ControlConnection.WriteMessage(websocket.PingMessage, []byte(clientID))
	}
}
func sendPump() {
	for {
		controlMessagBytes := <-sendChannel
		ControlConnection.WriteMessage(websocket.TextMessage, controlMessagBytes)
	}
}
func readPump() {
	for {
		mt, controlMessageBytes, err := ControlConnection.ReadMessage()
		if err != nil {
			log.Panic(err)
		}
		if mt == websocket.PingMessage {
			pingChannel <- true
			continue
		}
		if mt == websocket.TextMessage {
			log.Println("Received a message: ", string(controlMessageBytes))
			controlMessage, err := messages.ControlMessageBytesToControlMessage(controlMessageBytes)
			readChannel <- controlMessage
		}
	}

}
