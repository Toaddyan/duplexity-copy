package main

import (
	"context"
	"log"

	"github.com/duplexityio/duplexity/cmd/backend/pb"
	"google.golang.org/grpc"
)

func main() {
	// Create a connection to the gRPC server
	conn, err := grpc.Dial("localhost:9378", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer conn.Close()

	// Create a new Backend Client
	client := pb.NewBackendClient(conn)

	ctx := context.Background()
	response, err := client.GetConnection(ctx, &pb.GetConnectionRequest{
		RequestId: "random string",
		Hostname:  "todd.duplexity.io",
	})

	log.Println(response.Connection.GetHostname())
}
