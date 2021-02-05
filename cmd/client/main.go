package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/rancher/remotedialer"
)

func authorizer(protocol string, address string) bool {
	// this function should compare the protocol with the address
	return true
}

var clientID string

var config struct {
	WebsocketURI string `env:"WEBSOCKET_URI" envDefault:"ws://localhost:8080"`
	ClientID     string `env:"CLIENT_ID" envDefault:"client"`
}

func init() {
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", config)
}

func main() {
	ctx := context.Background()

	headers := http.Header{
		messages.ClientIDHeaderKey: []string{config.ClientID},
	}
	remotedialer.ClientConnect(ctx, fmt.Sprintf("%s/backend", config.WebsocketURI), headers, nil, authorizer, nil)
}
