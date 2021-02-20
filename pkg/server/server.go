package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/gorilla/websocket"
	"github.com/rancher/remotedialer"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Server is a remotedialer websocket that also contains a reverse proxy to each client
type Server struct {
	server     *remotedialer.Server
	Router     *router.Router
	Port       int
	WsHostName string
	hub        hub
}

// New returns a new Server
func New(router *router.Router, port int, wsHostname string, backendGrpcServer string) *Server {
	// Connect to backend service
	conn, err := grpc.Dial(backendGrpcServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Backend gRPC server failed to connect")
	}
	server := &Server{
		Router:     router,
		Port:       port,
		WsHostName: wsHostname,
		hub:        newHub(conn),
	}
	return server
}

// Serve HTTP for websocket server
func (s *Server) Serve() {
	s.server = remotedialer.New(s.authorizer, remotedialer.DefaultErrorWriter)
	// turn on the hub in background
	go s.hub.run()

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

func (s *Server) getClient(hostname string) (*Client, bool) {
	client, present := s.hub.clientMap[hostname]
	return client, present
}

func (s *Server) readControlInput(req *http.Request, ws *websocket.Conn) {
	for {
		_, rawMessage, err := ws.ReadMessage()
		log.Println("Received", string(rawMessage))
		if err != nil {
			log.Println("could not read: ", err)
		}
		cmdStruct := byteToStruct(rawMessage)
		log.Println("GOT A ", cmdStruct.Command)
		switch cmdStruct.Command {
		case messages.GetDataURI:
			s.respondDataURI(cmdStruct, ws)
		// case registerConnection:
		// 	s.responseConnection(req, cmdStruct, ws)
		case messages.Disconnect:
			s.disconnect(cmdStruct.ClientID, ws)
		case messages.UpdateTTL:
			log.Println("want to update ttl here")

		default:
			log.Println("unsupported command")
		}
	}
}

func (s *Server) controlHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running controlHandler")

	// websocket.upgrade into a websocket connection
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("can't upgrade to websocket")
	}
	defer ws.Close()

	s.readControlInput(req, ws)
}

// remotedialerHandler wraps around the remotedialer.Server.ServeHTTP
func (s *Server) remotedialerHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running remotedialerHandler")
	// Extract the clientID off of the incoming req
	hostname := req.Header.Get(messages.HostnameHeaderKey)
	log.Printf("Server.remotedialerHandler: got client id %s", hostname)
	// Send the request onto the remotedialer.Server
	s.server.ServeHTTP(w, req)

	// Delete the proxy that has corresponding clientID
	s.removeProxy(hostname)

}
