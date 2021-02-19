package server

type Client struct {
	hub      *hub
	hostname string
	// The websocket connection.
	conn connection

	// Buffered channel of outbound messages.
	send       chan message
	wsHostname string
}

func newClient(hostname string, hub hub, conn connection, wsHostname string) Client {
	return Client{
		hostname:   hostname,
		hub:        &hub,
		conn:       conn,
		send:       make(chan message),
		wsHostname: wsHostname,
	}
}
