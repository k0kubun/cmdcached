package server

import (
	"fmt"

	"github.com/sevlyar/go-daemon"
)

func Start() {
	context := &daemon.Context{
		PidFileName: "/tmp/cmdcached.pid",
		PidFilePerm: 0644,
		LogFileName: "/tmp/cmdcached.log",
		LogFilePerm: 0640,
		Umask:       027,
		Args:        []string{"cmdcached server"},
	}

	d, err := context.Reborn()
	if err != nil {
		fmt.Println(err)
		return
	}
	if d != nil {
		fmt.Println("cmdcached is successfully started")
		return
	}
	defer context.Release()

	server := NewServer()
	server.Run()
	server.Close()
}
