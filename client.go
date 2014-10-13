package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/k0kubun/cmdcached/server"
)

const (
	ClientSock = "/tmp/cmdcached.client.sock"
	maxBuf     = 1024 * 1024
)

func RequestCache() {
	os.Remove(ClientSock) // avoid "adress already in use"
	conn, err := net.DialUnix(
		"unix",
		&net.UnixAddr{ClientSock, "unix"},
		&net.UnixAddr{server.ServerSock, "unix"},
	)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	defer os.Remove(ClientSock)

	_, err = conn.Write([]byte("ghq list"))
	if err != nil {
		log.Println(err)
		return
	}

	var buf [maxBuf]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf(string(buf[:n]))
}
