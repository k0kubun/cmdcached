package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/k0kubun/cmdcached/server"
)

const (
	ClientSock = "/tmp/cmdcached.client.sock"
	maxBuf     = 1024 * 1024
)

type Client struct {
	conn *net.UnixConn
}

func NewClient() *Client {
	c := new(Client)

	os.Remove(ClientSock) // avoid "adress already in use"
	conn, err := net.DialUnix(
		server.ConnType,
		&net.UnixAddr{ClientSock, server.ConnType},
		&net.UnixAddr{server.ServerSock, server.ConnType},
	)
	if err != nil {
		fmt.Printf(err.Error())
	}
	c.conn = conn

	return c
}

func (c *Client) RequestCache(args []string) {
	if c.conn == nil {
		return
	}

	command := strings.Join(args, " ")
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		log.Println(err)
		return
	}

	var buf [maxBuf]byte
	n, err := c.conn.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf(string(buf[:n]))
}

func (c *Client) Close() {
	c.conn.Close()
	os.Remove(ClientSock)
}
