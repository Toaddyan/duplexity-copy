package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/websocket"
	"github.com/rancher/remotedialer"
)

// protobuff might be good here..

// Time constants
const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func listen() {
	for {
		_, rawMessage, err := controlConnection.ReadMessage()
		log.Println("RECEIVED A MESSAGE", string(rawMessage))
		if err != nil {
			log.Panic(err)
		}
		log.Println("convert byte to struct")
		cmdStruct := byteToStruct(rawMessage)
		log.Println("client ID: ", cmdStruct.ClientID, "and", config.ClientID)
		if cmdStruct.ClientID == config.ClientID {
			log.Println("clients match ")
			messagesChan <- cmdStruct
			log.Println("sent message")
		}
		log.Println("listen done ")

	}
}
func connect(ctx context.Context, resource string) {
	headers := http.Header{
		messages.HostnameHeaderKey: []string{config.ClientID},
		// Needs to be changed
		messages.ResourceHeaderKey: []string{resource},
	}
	remotedialer.ClientConnect(ctx, config.WebsocketURI, headers, nil, authorizer, nil)

}
func startConnection(ctx context.Context, resource string) {
	connect(ctx, resource)
}

func setWebsocketURI(cmd jsonCommand) error {
	log.Println("SETWEBSOCKET URI")
	if cmd.Command != messages.GetDataURI {
		return errors.New("unexpected servr Response")
	}
	if len(cmd.Args) == 0 {
		return errors.New("No args in Response")
	}

	config.WebsocketURI = cmd.Args[0]
	log.Println("Set WebsocketURI to ", config.WebsocketURI)
	return nil
}

// func write(command string, args ...string) {
// 	cmdStruct := newCommand(command)
// 	cmdStruct.args[0] = resource
// 	rawMessage := structToByte(cmdStruct)
// }

func writeControlMessage(cmd string, args ...string) {
	cmdStruct := newCommand(cmd)
	if args != nil {
		for _, arg := range args {
			cmdStruct.Args = append(cmdStruct.Args, arg)
		}
	}
	log.Println("cmdID ", cmdStruct.ClientID, " cmd", cmdStruct.Command)
	rawMessage := structToByte(cmdStruct)
	log.Println("SENDING MESSAGE:", string(rawMessage))
	controlConnection.WriteMessage(websocket.TextMessage, rawMessage)
}

func listenTerminate() {
	ticker := time.NewTicker((time.Minute))
	defer ticker.Stop()
	for {
		select {
		case msg := <-messagesChan:
			if msg.Command == messages.Disconnect {
				log.Println("messages.Disconnect successful")
				disconnectChan <- true
				return
			}
		case <-ticker.C:
			writeControlMessage(messages.UpdateTTL)

			// log.Println("Not a supported command %+v", msg)
		}
	}
}

// blocking
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			log.Println("terminating")
			return
		case <-disconnectChan:
			log.Println("disconnecting")
			return

		}
	}
}
