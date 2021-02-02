package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/rancher/remotedialer"
)

func authorizer(protocol string, address string) bool {
	// this function should compare the protocol with the address
	return true
}

var clientID string

func init() {
	flag.StringVar(&clientID, "clientid", "client", "Client ID")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	headers := http.Header{
		messages.ClientIDHeaderKey: []string{clientID},
	}
	remotedialer.ClientConnect(ctx, "ws://localhost:8080/backend", headers, nil, authorizer, nil)
}
