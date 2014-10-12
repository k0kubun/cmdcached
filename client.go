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

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Println(err)
		return
	}

	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s\n", string(buf[:n]))
}
