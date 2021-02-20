package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/cmd/backend/pb"
	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/websocket"
)

func (s *Server) setupBackEndConnection(clientID string, req *http.Request) {
	ctx := req.Context()
	resource := req.Header.Get(messages.ResourceHeaderKey)
	client := pb.NewBackendClient(s.hub.backendConnection)
	_, err := client.RegisterConnection(ctx, &pb.RegisterConnectionRequest{
		RequestId: "regster connection request id",
		Connection: &pb.Connection{
			UserId:    clientID,
			Hostname:  clientID,
			Websocket: s.WsHostName,
			Resource:  resource,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

// respondDataURI is the first request that happens
func (s *Server) respondDataURI(cmd jsonCommand, ws *websocket.Conn) {
	// Because this is the first time it connects, we need to register the client into the hub
	conn := newConnection(ws)
	// log.Println("SET UP NEW CONNECTION with client ", cmd.ClientID)
	client := newClient(cmd.ClientID, s.hub, *conn, s.WsHostName)
	// log.Println("REGISTERING CLIENT")
	s.hub.register <- &client
	// log.Println("WRITING MESAGE TO BACKEND")
	// client.wrapWriteControlMessage(messages.GetDataURI, fmt.Sprintf("ws://%s/backend", s.WsHostName))
	client.wrapWriteControlMessage(messages.GetDataURI, fmt.Sprintf("http://%s/backend", s.WsHostName))
	// log.Println("DONE with responsedata")
}

// func (s *Server) responseConnection(req *http.Request, cmd jsonCommand, ws *websocket.Conn) {
// 	_, present := s.getClient(cmd.ClientID)
// 	if present {
// 		log.Println("SETTING UP BACKEND CONN")
// 		s.setupBackEndConnection(cmd.ClientID, req)
// 		log.Println("WRITING BACK")
// 		writeControlMessage(ws, cmd.ClientID, cmd.Command)
// 	}
// }
func (s *Server) disconnect(clientID string, ws *websocket.Conn) {
	client, present := s.getClient(clientID)
	if !present {
		log.Panic("Cannot find client", clientID)
	}
	// Drop from remote dialer AND Redis
	s.hub.unregister <- client
	// really don't know if this is true
	s.server.RemovePeer(clientID)

	// writeMessage(ws, clientID, disconnect)
}
