package main

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/duplexityio/duplexity/pkg/server"
)

var config struct {
	HTTPPort          int    `env:"HTTP_PORT" envDefault:"8080"`
	Hostname          string `env:"WEBSOCKET_HOSTNAME"`
	BackendGrpcServer string `env:"BACKEND_GRPC_SERVER" envDefaut:"localhost:9378"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", config)

	if config.Hostname == "" {
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatalf("Could not determine hostname: %v\n", err)
		}
		config.Hostname = hostname
	}
}

func main() {
	// Forward-facing router
	router := router.New()

	// Websocket server
	server := server.New(router, config.HTTPPort, config.Hostname, config.BackendGrpcServer)
	server.Serve()
}
