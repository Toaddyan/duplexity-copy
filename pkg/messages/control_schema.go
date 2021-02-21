package messages

// ControlMessage is a wrapper for all control messages
type ControlMessage struct {
	// MessageType is the type of control message
	MessageType string `json:"message_type"`
	// Message is the JSON representation of the message
	Message string `json:"message"`
}

// Client is a Duplexity client
type Client struct {
	ID string `json:"id"`
}

// DiscoveryRequest is the first request the client when establishing a data plane connection
type DiscoveryRequest struct {
	Client *Client `json:"client"`
}

// DiscoveryResponse is the initial response from the server to establish the dataplen conection
type DiscoveryResponse struct {
	DataPlaneURI string `json:"data_plane_uri"`
}

// PipesRegistered is a response indicate succesful or failed registration of pipe(s)
type PipesRegistered struct {
	Succesful bool `json:"succesful"`
}
