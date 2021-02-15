package main

import (
	"encoding/json"

	"github.com/duplexityio/duplexity/cmd/backend/pb"
)

// TODO: Use the entireity of a Redis database, instead of a hash table
//       That way, we can set a TTL (expiration) on each entry
func hsetConnection(hostName string, connection *pb.Connection) (err error) {
	// Encode connection in JSON
	connectionJSON, err := json.Marshal(connection)
	if err != nil {
		return err
	}

	query := Redis.HSet("connections", hostName, string(connectionJSON))
	_, err = query.Result()
	if err != nil {
		return err
	}

	return
}

func hgetConnection(hostname string) (connection *Connection, err error) {
	connection = new(Connection)

	query := Redis.HGet("connections", hostname)
	connectionJSON, err := query.Result()
	if err != nil {
		return nil, buildNotFoundError("hgetConnection")
	}

	err = json.Unmarshal([]byte(connectionJSON), &connection)
	if err != nil {
		return nil, err
	}

	return
}
