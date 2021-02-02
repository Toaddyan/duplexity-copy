package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duplexityio/duplexity/pkg/router"
	"github.com/duplexityio/duplexity/pkg/server"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world!")

	// Instantiate server
	server := server.Server{}
	// Instantiate router
	router := router.Router{}
	// Serve HTTP for both server and router
	wsRouter := mux.NewRouter()
	wsRouter.Handle("/backend", server)
	wsRouter.Handle("/frontend", router)

	log.Println("Starting server")
	log.Fatalln(http.ListenAndServe(":8080", wsRouter))
}
