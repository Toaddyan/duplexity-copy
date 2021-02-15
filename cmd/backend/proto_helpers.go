package main

import "github.com/duplexityio/duplexity/cmd/backend/pb"

func buildConnectionResponse(connection *Connection) *pb.Connection {
	return &pb.Connection{
		UserId:    connection.UserID,
		Websocket: connection.Websocket,
		Hostname:  connection.HostName,
		Resource:  connection.Resource,
	}
}

func buildGetConnectionResponse(connection *pb.Connection) *pb.GetConnectionResponse {
	return &pb.GetConnectionResponse{
		Connection: connection,
	}
}
