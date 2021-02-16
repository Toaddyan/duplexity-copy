package router

import (
	"log"
	"sync"
	"time"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/rancher/remotedialer"
)

// Router is a controller to map
type Router struct {
	Proxies map[string]*httputil.ReverseProxy
	l       *sync.Mutex
}

// New ...
func New() *Router {
	router := &Router{
		Proxies: map[string]*httputil.ReverseProxy{},
		l:       &sync.Mutex{},
	}
	return router
}

// ServeHTTP starts the control service for websocket to hosts
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get the clientID and resource off the req
	hostname := req.Header.Get(messages.HostnameHeaderKey)
	if hostname == "" {
		log.Panicf("Router.ServeHTTP: clientID header key not provided")
	}

	proxy, present := r.CheckReverseProxies(hostname)
	if !present {
		return
	}
	// send request to user using ServeHTTP
	log.Printf("Router.ServeHTTP: Sending request to UserNode")
	proxy.ServeHTTP(w, req)
}

// CheckReverseProxies ensures that every host has one Proxy
func (r Router) CheckReverseProxies(hostname string) (*httputil.ReverseProxy, bool) {
	r.l.Lock()
	defer r.l.Unlock()
	proxy, present := r.Proxies[hostname]
	if !present {
		log.Println("Router.checkReverseProxies: No proxy in hostname:", hostname)
		return nil, false
	}
	return proxy, true
}

// GetProxy returns a Proxy
func (r Router) GetProxy(server *remotedialer.Server, hostname string) {
	_, present := r.CheckReverseProxies(hostname)
	if present {
		log.Fatal("Router.GetProxy: Proxy already exists")
	}
	dialer := server.Dialer(hostname, 15*time.Second)
	transport := &http.Transport{
		Dial: dialer,
	}
	reverseProxy := &httputil.ReverseProxy{
		Transport: transport,
		Director: func(req *http.Request) {
			resourceURL := req.Header.Get(messages.ResourceHeaderKey)
			log.Printf("Router.GetProxy: Got resourceURL: %s\n", resourceURL)

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
	log.Println("router.GetProxy: Building Proxy...")
	r.Proxies[hostname] = reverseProxy
}
