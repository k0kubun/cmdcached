package main

import "os"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		StartServer()
	} else {
		RequestCache()
	}
}
