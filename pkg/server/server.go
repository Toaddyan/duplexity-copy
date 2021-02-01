package server

import (
	"net/http"

	"github.com/complexityio/complexity/pkg/router"
	"github.com/rancher/remotedialer"
)

type Server struct {
	server remotedialer.Server
	router router.Router
}

// ServeHttp is a wrapper around remotedialer
func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Extract the clientID off of the incoming req

	// Send the request onto the remotedialer.Server
	s.server.ServeHTTP(w, req)

	// Delete the proxy that has corresponding clientID
	// s.router.RemoveProxy(clientID)
}

func (s Server) authorizer(req *http.Request) (clientID string, authed bool, err error) {
	// Extract the clientID off of the incoming req

	// Validate the client who is connecting
	// we make sure that they are actually who they claim they are

	// If they are not who they say they are, we return authed = false
	// and return err = "uh-oh, someone is not authenticated"

	// else

	// We do s.router.CreateProxy(s, clientID)

	// return authed = true

	// NOTE: Currently returning nothing, just so that go-lint does not complain
	return
}
