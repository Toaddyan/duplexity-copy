package main

import api "github.com/duplexityio/duplexity/cmd/backend/pb"

type server struct {
	api.UnimplementedBackendServer
}
