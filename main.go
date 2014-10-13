package main

import (
	"os"
	"runtime"

	"github.com/k0kubun/cmdcached/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	args := os.Args[1:]

	if len(args) == 0 {
		server.Start()
	} else {
		client := NewClient()
		client.RequestCache(args)
		client.Close()
	}
}
