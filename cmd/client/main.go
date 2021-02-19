package main

import (
	"context"
	"flag"
	"log"
	"net/url"
	"time"

	"github.com/caarlos0/env"
	"github.com/gorilla/websocket"
)

// TODO add resource key in flags
var config struct {
	WebsocketURI string `env:"WEBSOCKET_URI" envDefault:"ws://localhost:8080"`
	ClientID     string `env:"CLIENT_ID" envDefault:"client"`
}

var (
	controlConnection *websocket.Conn
	disconnectChan    chan bool
	messagesChan      chan jsonCommand
)

func authorizer(protocol, address string) bool {
	return true
}
func startControlConnection(websocketURI, clientID string) *websocket.Conn {
	url := url.URL{Scheme: "ws", Host: "websocket:8080", Path: "/control"}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("dial", err)
	}
	return conn
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
	ctx := context.Background()
	resource := flag.String("resource", "MISSING", "Application to be hauled")
	if *resource == "MISSING" {
		log.Fatal("need resource")
	}

	controlConnection = startControlConnection(config.WebsocketURI, config.ClientID)
	defer controlConnection.Close()
	// Listen for server Responses
	go listen()

	// Ensure connection at server is ok
	write(getDataURI)
	uriResponse := <-messagesChan
	err := setWebsocketURI(uriResponse)
	if err != nil {
		log.Fatal(err)
	}

	// Check if server is ok.
	write(registerConnection)
	connectionResponse := <-messagesChan
	go startConnection(ctx, *resource, connectionResponse)

	ticker := time.NewTicker((time.Minute))
	defer ticker.Stop()

	for {
		select {
		case <-disconnectChan:
			log.Println("Disconnect successful")
			return
		case <-ticker.C:
			write(updateTTL)
		default:
			log.Println("Not a supported command")
		}
	}

}
