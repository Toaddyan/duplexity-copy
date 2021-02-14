package main

import (
	"context"

	api "github.com/duplexityio/duplexity/cmd/backend/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) RegisterConnection(ctx context.Context, in *api.RegisterConnectionRequest) (response *emptypb.Empty, err error) {
	return nil, buildUnimplemntedError("RegisterConnection")
}

func (s *server) GetConnection(ctx context.Context, in *api.GetConnectionRequest) (response *api.GetConnectionResponse, err error) {
	return nil, buildUnimplemntedError("GetConnection")
}
