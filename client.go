package main

import (
	"log"
	"net"
	"os"
)

const (
	clientSock = "/tmp/cmdcached.client.sock"
)

func RequestCache() {
	conn, err := net.DialUnix(
		"unix",
		&net.UnixAddr{clientSock, "unix"},
		&net.UnixAddr{serverSock, "unix"},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(clientSock)

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
