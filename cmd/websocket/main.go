package main

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
)

var config struct {
	HTTPPort          int           `env:"HTTP_PORT" envDefault:"8080"`
	Hostname          string        `env:"WEBSOCKET_HOSTNAME"`
	BackendGrpcServer string        `env:"BACKEND_GRPC_SERVER" envDefaut:"localhost:9378"`
	WriteWait         time.Duration `env:"WRITE_WAIT" envDefault:"10s"`
	PongWait          time.Duration `env:"PONG_WAIT" envDefault:"60s"`
	PingPeriod        time.Duration `env:"PING_PERIOD" envDefault:"54s"`
	MaxMessageSize    int64         `env:"MAX_MESSAGE_SIZE" envDefault:"512"`
}

var (
	BackendGrpcConnection *grpc.ClientConn
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// Parse environment variables
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", config)

	// Determine hostname
	if config.Hostname == "" {
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatalf("Could not determine hostname: %v\n", err)
		}
		config.Hostname = hostname
	}

	// Dial BackendGrpcConnection
	BackendGrpcConnection, err = grpc.Dial(config.BackendGrpcServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to backend gRPC server: %v\n", err)
	}
}

func main() {
	// Forward-facing router
	// router := router.New()

	// // Websocket server
	// server := server.New(router, config.HTTPPort, config.Hostname, config.BackendGrpcServer)
	// server.Serve()
}
