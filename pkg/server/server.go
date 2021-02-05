package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/rancher/remotedialer"
)

// Server ...
type Server struct {
	server *remotedialer.Server
	Router *router.Router
	Port   int
}

// New returns a new server
func New(router *router.Router, port int) *Server {
	server := &Server{
		Router: router,
		Port:   port,
	}
	return server
}

// Serve ...
func (s *Server) Serve() {
	s.server = remotedialer.New(s.authorizer, remotedialer.DefaultErrorWriter)
	// httpRouter := mux.NewRouter()
	http.HandleFunc("/backend", s.backendHandler)
	http.HandleFunc("/frontend", s.Router.ServeHTTP)

	log.Println("Starting server")
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil))
}

func (s *Server) removeProxy(clientID string) {
	log.Println("removeProxy: removing proxy...", clientID)
	_, present := s.Router.CheckReverseProxies(clientID)
	if present {
		log.Printf("Removing %s from proxies\n", clientID)
		delete(s.Router.Proxies, clientID)
	}
	log.Panic("clientID not present")
}

// ServeHttp is a wrapper around remotedialer
func (s *Server) backendHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("running backendHandler")
	// Extract the clientID off of the incoming req
	clientID := req.Header.Get(messages.ClientIDHeaderKey)
	log.Printf("Server.backendHandler: got client id %s", clientID)
	// Send the request onto the remotedialer.Server
	s.server.ServeHTTP(w, req)
	// Delete the proxy that has corresponding clientID
	s.removeProxy(clientID)

}

func (s *Server) checkClientAuthentication(clientID string) (authed bool, err error) {
	if clientID == "" {
		authed = false
		err := fmt.Errorf("authorizer: missing %q header", messages.ClientIDHeaderKey)
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

func (s *Server) authorizer(req *http.Request) (clientID string, authed bool, err error) {
	// Extract the clientID off of the incoming req
	clientID = req.Header.Get(messages.ClientIDHeaderKey)
	log.Printf("Server.authorizer: Authorizing clientID: %s", clientID)
	// TODO ADD ERROR CHECKING
	authed, err = s.checkClientAuthentication(clientID)
	if !authed {
		return
	}
	log.Printf("Server.authorizer: Successful Authorization")
	s.Router.GetProxy(s.server, clientID)
	return
}
