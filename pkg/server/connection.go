package server

import "github.com/gorilla/websocket"

// encapsulate around client object 
type connection struct {
	send chan []byte
	ws   *websocket.Conn
}

func newConnection(ws *websocket.Conn) *connection {
	return &connection{
		send: make(chan []byte),
		ws:   ws,
	}
}
