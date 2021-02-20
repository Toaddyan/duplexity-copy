package server

type Client struct {
	hub      *hub
	clientID string
	// The websocket connection.
	conn connection

	// Buffered channel of outbound messages.
	send       chan jsonCommand
	wsHostname string
}

func (c *Client) sendPump() {
	for {
		select {
		case jsonCommand := <-c.send:
			c.wrapWriteControlMessage(jsonCommand.Command, jsonCommand.Args...)
		}
	}
}

func (c *Client) wrapWriteControlMessage(cmd string, args ...string) {
	c.conn.writeControlMessage(c.clientID, cmd, args...)
}

func newClient(hostname string, hub hub, conn connection, wsHostname string) Client {

	return Client{
		clientID:   hostname,
		hub:        &hub,
		conn:       conn,
		send:       make(chan jsonCommand),
		wsHostname: wsHostname,
	}
}
