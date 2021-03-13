package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/duplexityio/duplexity/pkg/router"
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
	// BackendGrpcConnection is a connection to go to the backend
	BackendGrpcConnection *grpc.ClientConn
)

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
			log.Fatalf("could not determine hostname: %v\n", err)
		}
		config.Hostname = hostname
		log.Println("set hostname to ", config.Hostname)
	}
	log.Println("Setting up BackendConnection")
	BackendGrpcConnection, err = grpc.Dial(config.BackendGrpcServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to backendgRPC server: %v\n", err)
	}
}
func main() {
	router := router.New()
	server := NewServer(router, config.HTTPPort, config.Hostname, BackendGrpcConnection)

	http.HandleFunc("/dial", server.remotedialerHandler)
	http.HandleFunc("/backend", server.backendHandler)
	http.HandleFunc("/frontend", server.Router.ServeHTTP)
	log.Println("Starting Websocket")
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.HTTPPort), nil))
}
