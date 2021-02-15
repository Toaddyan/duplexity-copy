package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/duplexityio/duplexity/cmd/backend/pb"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

var (
	// Redis is the client used to communicate withe Redis
	Redis *redis.Client
)

var config struct {
	BackendServerURI string `env:"BACKEND_SERVER_URI" envDefault:"0.0.0.0:9378"`
	RedisURI         string `env:"REDIS_URI" envDefault:"localhost:6379"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	listener, err := net.Listen("tcp", config.BackendServerURI)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Creating the gRPC server")
	grpcServer := grpc.NewServer()

	// Create new backendServer
	backendService := &service{}

	log.Println("Registering backend service to gRPC server")
	pb.RegisterBackendServer(grpcServer, backendService)

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
