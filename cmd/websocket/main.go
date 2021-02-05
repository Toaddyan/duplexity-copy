package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/duplexityio/duplexity/pkg/server"
)

var config struct {
	HTTPPort int `env:"HTTP_PORT" envDefault:"8080"`
}

func init() {
	log.SetFlags(log.Llongfile)

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", config)
}
func main() {
	fmt.Println("Hello world!")

	router := router.New()
	server := server.New(router, config.HTTPPort)
	server.Serve()

	// Serve HTTP for both server and router
	// wsRouter := mux.NewRouter()
	// wsRouter.Handle("/backend", server)
	// wsRouter.Handle("/frontend", router)

	// log.Println("Starting server")
	// log.Fatalln(http.ListenAndServe(":8080", wsRouter))
}
