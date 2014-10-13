package main

import (
	"os"

	"github.com/k0kubun/cmdcached/server"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		server.Start()
	} else {
		client := NewClient()
		client.RequestCache(args)
		client.Close()
	}
}
