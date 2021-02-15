package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/cmd/backend/pb"
	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/rancher/remotedialer"
	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	server            *remotedialer.Server
	Router            *router.Router
	Port              int
	Hostname          string
	backendConnection *grpc.ClientConn
}

// New returns a new server
func New(router *router.Router, port int, hostname string, backendGrpcServer string) *Server {
	conn, err := grpc.Dial(backendGrpcServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Backend gRPC server failed to connect")
	}
	server := &Server{
		Router:            router,
		Port:              port,
		Hostname:          hostname,
		backendConnection: conn,
	}
	return server
}

// Serve HTTP for websocket server
func (s *Server) Serve() {
	s.server = remotedialer.New(s.authorizer, remotedialer.DefaultErrorWriter)

	http.HandleFunc("/backend", s.remotedialerHandler)
	http.HandleFunc("/frontend", s.Router.ServeHTTP)
	http.HandleFunc("/control", s.controlHandler)

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

func (s *Server) controlHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running controlHandler")
	// TODO: Make it such that backendHandler does not begin the remotedialer WebSocket unless a valid control WebSocket
	//       is established.

}

// remotedialerHandler wraps around the remotedialer.Server.ServeHTTP
func (s *Server) remotedialerHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running remotedialerHandler")
	// Extract the clientID off of the incoming req
	hostname := req.Header.Get(messages.ClientIDHeaderKey)
	log.Printf("Server.remotedialerHandler: got client id %s", hostname)

	// TODO: Fix this hack
	ctx := req.Context()
	resource := req.Header.Get(messages.ResourceHeaderKey)
	client := pb.NewBackendClient(s.backendConnection)
	_, err := client.RegisterConnection(ctx, &pb.RegisterConnectionRequest{
		RequestId: "regster connection request id",
		Connection: &pb.Connection{
			UserId:    hostname,
			Hostname:  hostname,
			Websocket: s.Hostname,
			Resource:  resource,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Send the request onto the remotedialer.Server
	s.server.ServeHTTP(w, req)
	// Delete the proxy that has corresponding clientID
	s.removeProxy(hostname)
}

func (s *Server) checkClientAuthentication(clientID string) (authed bool, err error) {
	if clientID == "" {
		authed = false
		err := fmt.Errorf("authorizer: missing %q header", messages.ClientIDHeaderKey)
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
