package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	messagespb "github.com/duplexityio/duplexity/pkg/messages/pb"
	"github.com/gorilla/websocket"
)

var (
	// ControlConnection is websocket connection to the server for the control plane
	ControlConnection *websocket.Conn
	dataPlaneURI      string
	sendChannel       chan messagespb.ControlMessage
	readChannel       chan messagespb.ControlMessage
	pingChannel       chan bool
	disconnect        chan bool
)
var config struct {
	ControlWebsocketURI string `env:"WEBSOCKET_URI" envDefault:"ws://localhost:8080"`
	ClientID            string `env:"CLIENT_ID" envDefault:"client"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", config)

}

func main() {
	resource := *flag.String("resource", "http://hello", "Application to be hauled")
	sendChannel = make(chan messagespb.ControlMessage)
	readChannel = make(chan messagespb.ControlMessage)
	pingChannel = make(chan bool, 1)
	disconnect = make(chan bool, 1)

	log.Println("Starting ControlConnection")
	dialControlConnection(config.ControlWebsocketURI, config.ClientID)
	defer ControlConnection.Close()

	// ControlConnection Pumps
	go readPump()
	go writePump()
	go pingHandler(config.ClientID)

	// Sending DiscoveryRequest
	sendDiscoveryRequest(config.ClientID)
	dataPlaneURI, err := processDiscoveryResponse(config.ClientID)
	if err != nil {
		log.Fatal(err)
	}
	// Connecting to DataPlane using Remote Dialer for Proxy Service
	log.Println("Building RemoteDialer Pipe")
	go buildPipes(context.Background(), config.ClientID, dataPlaneURI, resource)

	// Listen For ControlMessages from server
	go listen()
	// Termination channels
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-c:
			log.Println("ctrl-c pressed, terminating")
		case <-disconnect:
			log.Println("Disconnecting")
			return
		}
	}
}
