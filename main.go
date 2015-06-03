package main

import (
	"flag"

	"github.com/marcboeker/gochat/client"
	"github.com/marcboeker/gochat/server"
)

var (
	mode     *string
	username *string
)

func init() {
	mode = flag.String("m", "client", "server or client")
	username = flag.String("u", "FooBar", "your username")
	flag.Parse()
}

func main() {
	if *mode == "server" {
		server.Start()
	} else {
		client.Start(*username)
	}
}
