package main

import (
	"context"
	"errors"
	"fmt"
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
		if err != nil {
			log.Panic(err)
		}
		cmdStruct := byteToStruct(rawMessage)

		if cmdStruct.clientID == config.ClientID {
			if cmdStruct.command == disconnect {
				disconnectChan <- true
			}
			messagesChan <- cmdStruct
		}
	}
}
func connect(ctx context.Context, resource string) {
	headers := http.Header{
		messages.HostnameHeaderKey: []string{config.ClientID},
		// Needs to be changed
		messages.ResourceHeaderKey: []string{resource},
	}
	remotedialer.ClientConnect(ctx, fmt.Sprintf("%s/backend", config.WebsocketURI), headers, nil, authorizer, nil)

}
func startConnection(ctx context.Context, resource string, cmd jsonCommand) error {
	if cmd.command != registerConnection {
		return errors.New("unexpected servr Response")
	}
	connect(ctx, resource)
	return nil
}

func setWebsocketURI(cmd jsonCommand) error {
	if cmd.command != getDataURI {
		return errors.New("unexpected servr Response")
	}
	if len(cmd.args) == 0 {
		return errors.New("No args in Response")
	}

	config.WebsocketURI = cmd.args[2]
	log.Println("Set WebsocketURI to ", config.WebsocketURI)
	return nil
}

// func write(command string, args ...string) {
// 	cmdStruct := newCommand(command)
// 	cmdStruct.args[0] = resource
// 	rawMessage := structToByte(cmdStruct)
// }

func write(cmd string, args ...string) {
	cmdStruct := newCommand(cmd)
	if args != nil {
		for _, arg := range args {
			cmdStruct.args = append(cmdStruct.args, arg)
		}
	}
	rawMessage := structToByte(cmdStruct)
	controlConnection.WriteMessage(websocket.TextMessage, rawMessage)
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Ctrl+C pressed: Disconnecting...")
		cmd := jsonCommand{
			clientID: config.ClientID,
			command:  disconnect,
		}
		rawMessage := structToByte(cmd)
		controlConnection.WriteMessage(websocket.TextMessage, rawMessage)
		return
	}()
}
