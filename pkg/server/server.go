package server

import (
	"fmt"
	"log"
	"net/http"

	g "github.com/duplexityio/duplexity/pkg/messages"
	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/rancher/remotedialer"
)

// Server ...
type Server struct {
	server remotedialer.Server
	router *router.Router
}

func (s Server) removeProxy(clientID string) {
	_, present := s.router.GetProxy(clientID)
	if present {
		log.Printf("Removing %s from proxies\n", clientID)
		delete(s.router.Proxies, clientID)
	}
	log.Panic("clientID not present")
}

// ServeHttp is a wrapper around remotedialer
func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Extract the clientID off of the incoming req
	clientID := req.Header.Get(g.ClientIDHeaderKey)
	// Send the request onto the remotedialer.Server
	s.server.ServeHTTP(w, req)
	// Delete the proxy that has corresponding clientID
	s.removeProxy(clientID)

}

func (s Server) checkClient(clientID string) (authed bool, err error) {
	if clientID == "" {
		authed = false
		err := fmt.Errorf("authorizer: missing %q header", g.ClientIDHeaderKey)
		// err = errors.New(Sprintf("authorizer: missing clientID header")
		return authed, err
	}

	// TODO: Use OAuth
	// Validate the client who is connecting
	// we make sure that they are actually who they claim they are
	// If they are not who they say they are, we return authed = false
	// and return err = "uh-oh, someone is not authenticated"
	return true, nil
}

func (s Server) authorizer(req *http.Request) (clientID string, authed bool, err error) {
	// Extract the clientID off of the incoming req
	clientID = req.Header.Get(g.ClientIDHeaderKey)

	authed, err = s.checkClient(clientID)
	if !authed {
		return
	}
	s.router.CreateProxy(&s.server, clientID)
	return
}
