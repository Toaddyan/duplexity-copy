package main

import (
	"log"

	messagespb "github.com/duplexityio/duplexity/pkg/messages/pb"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func pingHandler() {
	for {
		<-pingChannel
		ControlConnection.WriteMessage(websocket.PingMessage, []byte(config.ClientID))
	}
}

func writePump() {
	for {
		controlMessage := <-sendChannel
		log.Printf("writePump has received a message: %+v", controlMessage.ClientID)
		controlMessageBytes, err := proto.Marshal(controlMessage)
		if err != nil {
			log.Panicf("%v\n", err)
		}
		ControlConnection.WriteMessage(websocket.TextMessage, controlMessageBytes)
		log.Println("WritePump has sent message")
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
		log.Println("received back a message")
		if mt == websocket.PingMessage {
			pingChannel <- true
		}

		if mt == websocket.TextMessage {
			controlMessage := messagespb.ControlMessage{}
			err := proto.Unmarshal(controlMessageBytes, &controlMessage)
			if err != nil {
				log.Panicf("%v\n", err)
			}
			log.Println("got control mesage in read pump: ", controlMessage.GetClientID())
			readChannel <- &controlMessage
		}
	}
}
