package server

import (
	"sync"

	"google.golang.org/grpc"
)

type subscription struct {
	conn     *connection
	hostname string
}

type message struct {
	data     []byte
	clientID string
}

type hub struct {
	l                 sync.Mutex
	backendConnection *grpc.ClientConn
	clientMap         map[string]*Client
	register          chan *Client
	unregister        chan *Client
	broadcast         chan message
}

func newHub(backendConnection *grpc.ClientConn) hub {
	return hub{
		backendConnection: backendConnection,
		clientMap:         make(map[string]*Client),
		register:          make(chan *Client),
		unregister:        make(chan *Client),
		broadcast:         make(chan message),
	}
}

//use function calls in case statements
func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			h.l.Lock()
			h.clientMap[client.clientID] = client
			h.l.Unlock()
			go client.sendPump()

		case client := <-h.unregister:
			_, present := h.clientMap[client.clientID]
			if present {
				delete(h.clientMap, client.clientID)
				close(client.send)
			}
			// case message := <-h.broadcast:
			// 	for _, client := range h.clientMap {
			// 		client.send <- message

			// 	}
		}
	}
}
