package router

import (
	"log"
	"sync"
	"time"

	"net/http"
	"net/http/httputil"
	"net/url"

	g "github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/mux"
	"github.com/rancher/remotedialer"
)

// TODO: Test ReverseProxys for race conditions

// Router ...
type Router struct {
	Proxies map[string]*httputil.ReverseProxy
	l       *sync.Mutex
}

// GetProxy returns a proxy
func (r Router) GetProxy(clientID string) (*httputil.ReverseProxy, bool) {
	r.l.Lock()
	defer r.l.Unlock()
	proxy, present := r.Proxies[clientID]
	if !present {
		log.Panic("no proxy in clientID:", clientID)
		return nil, false
	}
	return proxy, true
}

// ServeHTTP ...
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get the clientID and resource off the req
	clientIDs, present := req.Header[g.ClientIDHeaderKey]
	if !present {
		log.Printf("No clientID in the header")
		return
	}
	clientID := clientIDs[0]
	proxy, present := r.GetProxy(clientID)
	if !present {
		return
	}
	// send request to UserNode using ServeHTTP
	proxy.ServeHTTP(w, req)
}

// CreateProxy ...
func (r Router) CreateProxy(server *remotedialer.Server, clientID string) {
	_, present := r.GetProxy(clientID)
	if present {
		log.Fatal("Can't create a proxy which already exists...")
	}
	dialer := server.Dialer(clientID, 15*time.Second)
	transport := &http.Transport{
		Dial: dialer,
	}
	reverseProxy := &httputil.ReverseProxy{
		Transport: transport,
		Director: func(req *http.Request) {

			// extract the resource out of the headers

			// DO NOT do this
			vars := mux.Vars(req)

			// get the resource off the headers, DO NOT do this
			resourceURL, present := vars["resource"]
			if !present {
				log.Panicln("No resource provided")
			}

			log.Printf("Got resourceURL: %s\n", resourceURL)

			resource, err := url.Parse(resourceURL)
			if err != nil {
				log.Panicln(err)
			}

			// TODO: Something wasn't working right with the req.URL, the path of /frontend was getting forwarded along
			//       so as a really quick and dirty hack I am simply setting req.URL to the same as resource
			//       This may be a bad idea?
			//       It is probably related to this issue:
			//       https://github.com/golang/go/issues/28168
			//       req.URL.Path = resource.Path probably fixes it?
			// req.URL.Scheme = resource.Scheme
			// req.URL.Host = resource.Host
			req.URL = resource

			req.Host = resource.Host
		},
	}
	r.Proxies[clientID] = reverseProxy
}
