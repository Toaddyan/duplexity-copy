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

func (s *Server) respondDataURI(cmd jsonCommand, ws *websocket.Conn, req *http.Request) {
	s.l.Lock()
	defer s.l.Unlock()
	conn := newConnection(ws)
	client := newClient(cmd.clientID, s.hub, *conn, s.WsHostName)
	s.hub.register <- &client
	s.setupBackEndConnection(cmd.clientID, req)
	write(ws, cmd.clientID, getDataURI, s.WsHostName)

}

func (s *Server) responseConnection(cmd jsonCommand, ws *websocket.Conn) {
	s.l.Lock()
	defer s.l.Unlock()
	client, present := s.getClient(cmd.clientID)
	if present {
		connMsg := fmt.Sprintf("%s %s", client.hostname, registerConnection)
		ws.WriteMessage(websocket.TextMessage, []byte(connMsg))
	}
}
func (s *Server) disconnect(clientID string, ws *websocket.Conn) {
	s.l.Lock()
	defer s.l.Unlock()
	client, present := s.getClient(clientID)
	if present {
		// Drop from remote dialer AND Redis
		s.hub.unregister <- client
		// really don't know if this is true
		s.server.RemovePeer(clientID)
		// Need to clean up database..
		s.hub.backendConnection.Close()
	}
	write(ws, clientID, disconnect)
}
