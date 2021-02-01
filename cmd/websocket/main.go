package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world!")
	// Instantiate server
	// Instantiate router

	// Serve HTTP for both server and router
	wsRouter := mux.NewRouter()
	wsRouter.Handle("/backend", server)
	wsRouter.Handle("/frontend", router)

	log.Println("Starting server")
	log.Fatalln(http.ListenAndServe(":8080", wsRouter))
}
