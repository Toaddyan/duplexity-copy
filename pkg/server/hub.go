package server

import (
	"google.golang.org/grpc"
)

type subscription struct {
	conn     *connection
	hostname string
}

type message struct {
	data     []byte
	hostname string
}

type hub struct {
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
		case c := <-h.register:
			h.clientMap[c.hostname] = c

		case c := <-h.unregister:
			_, present := h.clientMap[c.hostname]
			if present {
				delete(h.clientMap, c.hostname)
				close(c.send)
			}
		case m := <-h.broadcast:
			for _, client := range h.clientMap {
				client.send <- m

			}
		}
	}
}
