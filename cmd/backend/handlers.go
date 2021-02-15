package main

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/duplexityio/duplexity/cmd/backend/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) RegisterConnection(ctx context.Context, in *pb.RegisterConnectionRequest) (response *emptypb.Empty, err error) {
	log.Printf("Starting request for RegisterConnection: %v\n", in)

	defer func() {
		if r := recover(); r != nil {
			err = buildInternalError("RegisterConnection", r, string(debug.Stack()))
			log.Println(err)
		}
	}()

	err = hsetConnection(in.Connection.GetHostname(), in.GetConnection())
	if err != nil {
		logInternalError(err)
		return nil, err
	}
	response = &emptypb.Empty{}

	log.Println("Request finished")
	return
}

func (s *service) GetConnection(ctx context.Context, in *pb.GetConnectionRequest) (response *pb.GetConnectionResponse, err error) {
	log.Printf("Starting request for GetConnection: %v\n", in)

	defer func() {
		if r := recover(); r != nil {
			err = buildInternalError("GetConnection", r, string(debug.Stack()))
			log.Println(err)
		}
	}()

	connection, err := hgetConnection(in.GetHostname())
	if err != nil {
		logInternalError(err)
		return nil, err
	}

	connectionResponse := buildConnectionResponse(connection)
	response = buildGetConnectionResponse(connectionResponse)

	log.Println("Request finished")
	return
}
