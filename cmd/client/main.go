package main

import (
	"log"

	"github.com/duplexityio/duplexity/pkg/proxyserver"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	proxy := proxyserver.ProxyServer{Port: 9999}

	proxy.Serve()
}
