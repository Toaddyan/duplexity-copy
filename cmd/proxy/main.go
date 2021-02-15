package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/duplexityio/duplexity/pkg/proxy"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

var config struct {
	BackendGrpcServer string `env:"BACKEND_GRPC_SERVER" envDefault:"localhost:9378"`
	ProxyHTTPPort     int    `env:"PROXY_HTTP_PORT" envDefault:"9090"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("Could not parse config: %v\n", err)
	}
}

func main() {
	proxy := proxy.New(config.ProxyHTTPPort, config.BackendGrpcServer)
	proxy.Serve()
}
