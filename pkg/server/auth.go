package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/pkg/messages"
)

func (s *Server) checkClientAuthentication(clientID string) (authed bool, err error) {
	if clientID == "" {
		authed = false
		err := fmt.Errorf("authorizer: missing %q header", messages.HostnameHeaderKey)
		// err = errors.New(Sprintf("authorizer: missing clientID header")
		return authed, err
	}
	// TODO: Use OAuth

	// cfg := oauth.NewClientConfig("https://accounts.google.com")

	// http.HandleFunc("/", cfg.LoginUser)
	// http.HandleFunc("/auth/google/callback", cfg.CallbackHandler)

	// log.Printf("listening on http://%s/", "127.0.0.1:5556")
	// log.Fatal(http.ListenAndServe("127.0.0.1:5556", nil))

	// Validate the client who is connecting
	// we make sure that they are actually who they claim they are
	// If they are not who they say they are, we return authed = false
	// and return err = "uh-oh, someone is not authenticated"
	return true, nil
}

func (s *Server) authorizer(req *http.Request) (clientID string, authed bool, err error) {
	// Extract the clientID off of the incoming req
	clientID = req.Header.Get(messages.HostnameHeaderKey)
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
