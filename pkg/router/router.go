package router

import (
	"log"
	"sync"
	"time"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/rancher/remotedialer"
)

// TODO: Test ReverseProxys for race conditions

// Router is
type Router struct {
	proxies map[string]*httputil.ReverseProxy
	l       sync.Mutex
}

func (r Router) getProxy(clientID string) *httputil.ReverseProxy {
	r.l.Lock()
	defer r.l.Unlock()

	// TODO: Return the correct reverseproxy, this is just to stop go-lint from complaining
	return &httputil.ReverseProxy{}
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get the clientID and resource off the req

	// grab proxy from proxies map
	//  proxy = r.getProxy(clientID)

	// send request to UserNode using ServeHTTP
}

func (r Router) CreateProxy(server *remotedialer.Server, clientID string) {
	// Check if proxy already exists
	// call getProxy, if it returns... then something has seriously gone wrong

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
	r.proxies[clientID] = reverseProxy
}
