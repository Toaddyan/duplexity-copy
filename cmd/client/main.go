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
// TODO: 
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
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", config)
}

func main() {
	// Redirect the user to oauth service.
	//  wait until user autenticates

	// Check -> Am I authenticated...

	//  in linux machine... it goes into ~/.config
	//  Be able to look for this in platform agnostic way
	//  hydrate a jwt and dehydrate a JWT  encode decode ?
	ctx := context.Background()

	headers := http.Header{
		// add to the headers... here are the credentials.
		messages.HostnameHeaderKey: []string{config.ClientID},
		messages.ResourceHeaderKey: []string{"http://hello"},
	}
	remotedialer.ClientConnect(ctx, fmt.Sprintf("%s/backend", config.WebsocketURI), headers, nil, authorizer, nil)
}
