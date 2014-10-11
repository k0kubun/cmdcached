package main

import (
	"fmt"

	"github.com/sevlyar/go-daemon"
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
	for {
	}
}
