package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/howeyc/fsnotify"
	"github.com/sevlyar/go-daemon"
)

const (
	serverSock = "/tmp/cmdcached.sock"
)

func StartServer() {
	context := &daemon.Context{
		PidFileName: "/tmp/cmdcached.pid",
		PidFilePerm: 0644,
		LogFileName: "/tmp/cmdcached.log",
		LogFilePerm: 0640,
		WorkDir:     "/tmp",
		Umask:       027,
		Args:        []string{"cmdcached server"},
	}

	d, err := context.Reborn()
	if err != nil {
		fmt.Println("cmdcached is already started")
		return
	}
	if d != nil {
		fmt.Println("cmdcached is successfully started")
		return
	}
	defer context.Release()

	server := new(Server)
	server.Run()
}

type Server struct {
}

func (s *Server) Run() {
	go s.Watch()

	l, err := net.ListenUnix(
		"unix",
		&net.UnixAddr{serverSock, "unix"},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(serverSock)

	for {
		conn, err := l.AcceptUnix()
		if err != nil {
			log.Fatal(err)
		}

		go s.Serve(conn)
	}
}

func (s *Server) Serve(conn *net.UnixConn) {
	defer conn.Close()

	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", string(buf[:n]))
}

func (s *Server) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch("/Users/k0kubun/src")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
			log.Println("[Event]", ev)
		case err = <-watcher.Error:
			log.Println("[Error]", err)
		}
	}
}
