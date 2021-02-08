package main

import (
	"log"

	"github.com/duplexityio/duplexity/pkg/proxy"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	proxy := proxy.New(9090)

	proxy.Serve()
}