package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/mux"
)

// Proxy ...
type Proxy struct {
	Port         int
	reverseproxy *httputil.ReverseProxy
}

// New creates proxy struct
func New(port int) *Proxy {
	proxy := &Proxy{
		Port: port,
	}

	proxy.reverseproxy = &httputil.ReverseProxy{
		Director: proxy.director,
	}
	return proxy
}

// Serve will start the proxy service
func (proxy Proxy) Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/", proxy.reverseproxy.ServeHTTP)
	log.Printf("Serving HTTP on %d\n", proxy.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", proxy.Port), router)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Bye bye")
}

func (proxy Proxy) director(req *http.Request) {
	log.Printf("New req: %#v\n", req)

	// host := req.Host
	// clientID.duplexity.io
	clientID := strings.Split(req.Host, ".")[0]
	log.Println("director: clientID: ", clientID)
	req.Header.Set(messages.ClientIDHeaderKey, clientID)

	// lookup on the host
	// TODO: Do a lookup on the backend, to figure out what websocket the user node is connected to
	// req.Header.Set(messages.ResourceHeaderKey, "http://localhost:1234")
	req.Header.Set(messages.ResourceHeaderKey, "http://ipinfo.io")
	// THIS SHOULD BE A LOOK UP INTO THE DATABASE FOR WHAT THE USER WANTS TO EXPOSE
	

	resource, err := url.Parse("http://websocket:8080/frontend")
	if err != nil {
		log.Panicln("director: can't parse url websocket/frontend ")
	}

	// req.URL.Scheme = resource.Scheme
	// req.URL.Host = resource.Host
	req.URL = resource
	req.Host = resource.Host

}

// docker run -d -p 7070:8080 --rm -t mendhak/http-https-echo:17
