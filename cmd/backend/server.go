package main

import api "github.com/duplexityio/duplexity/cmd/backend/pb"

type service struct {
	api.UnimplementedBackendServer
}
