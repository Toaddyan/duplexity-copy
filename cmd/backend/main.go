package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	api "github.com/duplexityio/duplexity/cmd/backend/pb"
	"google.golang.org/grpc"
)

var config struct {
	BackendServerURI string `env:"BACKEND_SERVER_URI" envDefault:"0.0.0.0:9378"`
}

func init() {
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", config.BackendServerURI)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Creating the gRPC server")
	grpcServer := grpc.NewServer()

	// Create new backendServer
	backendServer := &server{}

	log.Println("Registering Backend server to gRPC server")
	api.RegisterBackendServer(grpcServer, backendServer)

	go func() {
		log.Println("Opening gRPC server to world")
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	<-sig
	log.Println("Received SIGINT")
	log.Println("Gracefully stopping grpcServer")
	grpcServer.GracefulStop()

	log.Println("Bye bye")
}
