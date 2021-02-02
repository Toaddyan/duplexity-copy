package proxyserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ProxyServer is an HTTP proxy server
type ProxyServer struct {
	Port int
}

// Serve starts the proxy server
func (server *ProxyServer) Serve() {
	router := mux.NewRouter()

	// Register HTTP handlers
	router.HandleFunc("/health", server.healthHandler)
	router.HandleFunc("/proxy", server.proxyHandler)

	http.Handle("/", router)

	log.Println("Starting proxyserver")
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("Bye bye")
}

func (server *ProxyServer) healthHandler(w http.ResponseWriter, req *http.Request) {
	message, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		"healthy",
	})
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(message)
}

func copyHeaders(destination, source http.Header) {
	for key, values := range source {
		for _, value := range values {
			destination.Add(key, value)
		}
	}
}

func (server *ProxyServer) proxyHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Method: %s, url: %v", req.Method, req.URL)
	if req.Method == http.MethodGet {

		log.Printf("Header: %#v", req.Header)

		resourceHeaders, present := req.Header["Proxyserver-Resource"]
		if !present { // case where there are no headers
			log.Panicln("Resource not present")
		} else if len(resourceHeaders) > 1 { // case where there are multiple headers
			log.Panicln("Too many resources provided")
		}
		resource := resourceHeaders[0]

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Minute,
		}

		proxyReq, err := http.NewRequest("GET", resource, nil)
		if err != nil {
			log.Panicln(err)
		}

		copyHeaders(proxyReq.Header, req.Header)

		log.Printf("Making GET request to %s\n", resource)
		resp, err := client.Do(proxyReq)
		if err != nil {
			// TODO: Do something special if the request times out, which can be retrieved from err
			log.Panicln(err)
		}

		copyHeaders(w.Header(), resp.Header)

		log.Printf("Received StatusCode: %d\n", resp.StatusCode)
		w.WriteHeader(resp.StatusCode)

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 { // case where resp was HTTP 2xx
			written, err := io.Copy(w, resp.Body)
			if err != nil {
				log.Panicln(err)
			}
			log.Printf("Copied %d bytes into response\n", written)
		}

	} else {
		http.Error(w, "GET is the only supported method", http.StatusMethodNotAllowed)
	}
}
