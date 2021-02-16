package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/duplexityio/duplexity/cmd/backend/pb"
	"github.com/duplexityio/duplexity/pkg/messages"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Proxy is a reverse proxy responsible for finding where our client is. 
type Proxy struct {
	Port              int
	reverseproxy      *httputil.ReverseProxy
	backendConnection *grpc.ClientConn
}

// New creates new Proxy
func New(port int, backendGrpcServer string) *Proxy {
	log.Println("Creating a new Proxy")
	// TODO: Fatal if doesn't finish within 15 seconds
	conn, err := grpc.Dial(backendGrpcServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Backend gRPC server failed to connect")
	}

	log.Println("heyo")

	proxy := &Proxy{
		Port:              port,
		backendConnection: conn,
	}

	proxy.reverseproxy = &httputil.ReverseProxy{
		Director: proxy.director,
	}
	return proxy
}

// Serve starts the proxy service
func (proxy Proxy) Serve() {
	log.Println("Starting Proxy.Serve")

	defer proxy.backendConnection.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", proxy.reverseproxy.ServeHTTP)

	log.Printf("Serving HTTP on %d\n", proxy.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", proxy.Port), router)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Bye bye")

}
// director is a custom controller for the reverse proxy
func (proxy Proxy) director(req *http.Request) {
	log.Printf("New req: %#v\n", req)

	ctx := req.Context()

	hostname := strings.Split(req.Host, ".")[0]
	log.Println("director: clientID: ", hostname)
	req.Header.Set(messages.HostnameHeaderKey, hostname)

	// Create a new backend client
	client := pb.NewBackendClient(proxy.backendConnection)
	// Do a getConnection request on Backend service
	response, err := client.GetConnection(ctx, &pb.GetConnectionRequest{
		RequestId: "SomeGetConnectionRequest",
		Hostname:  hostname,
	})
	if err != nil {
		//TODO: make this 404
		log.Panicln("Couldn't get valid Connection Request")
	}

	req.Header.Set(messages.ResourceHeaderKey, response.Connection.GetResource())
	resource, err := url.Parse(fmt.Sprintf("http://%s:8080/frontend", response.Connection.GetWebsocket()))
	if err != nil {
		log.Panicln("director: can't parse url websocket/frontend")
	}

	// req.URL.Scheme = resource.Scheme
	// req.URL.Host = resource.Host
	req.URL = resource
	req.Host = resource.Host

}

// docker run -d -p 7070:8080 --rm -t mendhak/http-https-echo:17
