package server

import (
	"log"
	"net"
	"os"

	"github.com/howeyc/fsnotify"
)

const (
	ServerSock = "/tmp/cmdcached.sock"
)

type Server struct {
	config *Config
}

func (s *Server) Run() {
	s.config = new(Config)
	s.config.Load()

	go s.Watch()

	os.Remove(ServerSock) // avoid "address already in use"
	l, err := net.ListenUnix(
		"unix",
		&net.UnixAddr{ServerSock, "unix"},
	)

	if err != nil {
		log.Println(err)
		return
	}
	defer os.Remove(ServerSock)

	for {
		conn, err := l.AcceptUnix()
		if err != nil {
			log.Println(err)
			return
		}

		go s.Serve(conn)
	}
}

func (s *Server) Serve(conn *net.UnixConn) {
	defer conn.Close()

	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}
	result, err := CachedExec(string(buf[:n]))
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	conn.Write([]byte(result))
}

func (s *Server) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return
	}

	err = watcher.Watch("/Users/k0kubun/src")
	if err != nil {
		log.Println(err)
		return
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
