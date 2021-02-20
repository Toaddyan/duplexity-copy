package server

import "github.com/gorilla/websocket"

// encapsulate around client object
type connection struct {
	send chan []byte
	ws   *websocket.Conn
}

func (c *connection) writeControlMessage(clientID, cmd string, args ...string) {
	cmdStruct := newCommand(cmd)
	cmdStruct.ClientID = clientID
	if args != nil {
		for _, arg := range args {
			cmdStruct.Args = append(cmdStruct.Args, arg)
		}
	}
	rawMessage := structToByte(cmdStruct)
	c.ws.WriteMessage(websocket.TextMessage, rawMessage)
}

func newConnection(ws *websocket.Conn) *connection {
	return &connection{
		send: make(chan []byte),
		ws:   ws,
	}
}
