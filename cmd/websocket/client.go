package main

import (
	"log"
	"time"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/websocket"
)

// Client is a representation of a connected Duplexity client
type Client struct {
	ID         string
	Connection *websocket.Conn
	Hub        *Hub

	send chan messages.ControlMessage
}

func (client *Client) handleMessage(message interface{}) {
	log.Printf("Received message: %v\n", message)
}

func (client *Client) readPump() {
	defer func() {
		client.Hub.Unregister <- client
		client.Connection.Close()
	}()

	client.Connection.SetReadLimit(config.MaxMessageSize)
	client.Connection.SetReadDeadline(time.Now().Add(config.PongWait))
	client.Connection.SetPongHandler(func(string) error {
		client.Connection.SetReadDeadline(time.Now().Add(config.PongWait))
		return nil
	})

	for {
		messageType, messageBytes, err := client.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Could not read for %s, receeived close error: %v\n", client.ID, err)
			}
			break
		}
		if messageType == websocket.BinaryMessage {
			message, err := messages.ControlMessageBytesToMessage(messageBytes)
			if err != nil {
				log.Printf("Could not decode messageBytes: %v\n", err)
			}
			client.handleMessage(message)
		}
	}

}
func (client *Client) writePump() {
	ticker := time.NewTicker(config.PingPeriod)
	defer func() {
		ticker.Stop()
		client.Connection.Close()
	}()

	for {
		select {
		case controlMessage, ok := <-client.send:
			if !ok { // case where send channel was closed
				log.Printf("The send channel for %s was closed; now writing CloseMessage\n", client.ID)
				client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := client.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("Could not get a writer for %s: %v\n", client.ID, err)
				return
			}

			controlMessageBytes, err := messages.ControlMessageToControlMessageBytes(controlMessage)
			if err != nil {
				log.Printf("Could not get bytes version of control message: %v\n", err)
				return
			}

			_, err = writer.Write(controlMessageBytes)
			if err != nil {
				log.Printf("Could not write for %s: %v\n", client.ID, err)
				return
			}

			err = writer.Close()
			if err != nil {
				log.Printf("Could not close writer for %s: %v\n", client.ID, err)
				return
			}

		case <-ticker.C:
			client.Connection.SetWriteDeadline(time.Now().Add(config.WriteWait))
			err := client.Connection.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Printf("Could not write ping for %s: %v\n", client.ID, err)
				return
			}
		}

	}

}
