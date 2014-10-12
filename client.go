package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	clientSock = "/tmp/cmdcached.client.sock"
	maxBuf     = 1024 * 1024
)

func RequestCache() {
	os.Remove(clientSock) // avoid "adress already in use"
	conn, err := net.DialUnix(
		"unix",
		&net.UnixAddr{clientSock, "unix"},
		&net.UnixAddr{serverSock, "unix"},
	)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	defer os.Remove(clientSock)

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
