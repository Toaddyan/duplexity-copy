package main

import (
	"fmt"
	"log"

	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/duplexityio/duplexity/pkg/server"
)

func init() {
	log.SetFlags(log.Llongfile)
}
func main() {
	fmt.Println("Hello world!")

	// implement a server.serve
	//   does lines 30-35
	// Instantiate server
	router := router.New()
	server := server.New(router, 8080)
	server.Serve()
	// Instantiate router

	// Serve HTTP for both server and router
	// wsRouter := mux.NewRouter()
	// wsRouter.Handle("/backend", server)
	// wsRouter.Handle("/frontend", router)

	// log.Println("Starting server")
	// log.Fatalln(http.ListenAndServe(":8080", wsRouter))
}
