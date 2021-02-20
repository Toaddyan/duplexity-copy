package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/duplexityio/duplexity/pkg/messages"
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
	dataURI           string
)

func authorizer(protocol, address string) bool {
	return true
}
func startControlConnection(websocketURI, clientID string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("%s/control", websocketURI), nil)
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
	log.Println("starting client")
	ctx := context.Background()
	log.Println("huh. we here yet?")
	resource := flag.String("resource", "http://hello", "Application to be hauled")
	if *resource == "" {
		log.Fatalln("need resource")
	}

	controlConnection = startControlConnection(config.WebsocketURI, config.ClientID)
	defer controlConnection.Close()
	// Listen for server Responses
	messagesChan = make(chan jsonCommand)
	go listen()

	log.Println("ENSURE CONNECTION AT SERVER AND SET BACKEND")
	// Ensure connection at server is ok
	writeControlMessage(messages.GetDataURI)
	log.Println("WAITING FOR MESSAGE FROM SERVER")
	uriResponse := <-messagesChan
	log.Println("RECEIVED MESSAGE SETTING WEBSCOKET URI")
	err := setWebsocketURI(uriResponse)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("CHECK IF CONNECTION IS THERE. ")
	// Check if server is ok.
	// writeControlMessage(registerConnection)
	log.Println("STARTING REMOTE DIALER CONNECTION ")
	go startConnection(ctx, *resource)

	log.Println("Waiting for response from websocket")
	connectionResponse := <-messagesChan
	log.Printf("Got back response from websocket: %s\n", connectionResponse)

	go listenTerminate()
	setupCloseHandler()

}

// WHY IS THIS NOWF WJATTTT
