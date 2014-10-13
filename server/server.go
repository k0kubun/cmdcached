package server

import (
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/howeyc/fsnotify"
)

const (
	ServerSock = "/tmp/cmdcached.sock"
	ConnType   = "unix"
)

type Server struct {
	config      *Config
	resultCache map[string]string
	listener    *net.UnixListener
}

func NewServer() *Server {
	s := new(Server)
	s.config = new(Config)
	s.config.Load()
	s.resultCache = make(map[string]string)

	os.Remove(ServerSock) // avoid "address already in use"
	l, err := net.ListenUnix(
		ConnType,
		&net.UnixAddr{ServerSock, ConnType},
	)
	if err != nil {
		log.Println(err)
	}
	s.listener = l

	return s
}

func (s *Server) Run() {
	for {
		conn, err := s.listener.AcceptUnix()
		if err != nil {
			log.Println(err)
			continue
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
	result, err := s.cachedExec(string(buf[:n]))
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	conn.Write([]byte(result))
}

func (s *Server) Close() {
	s.listener.Close()
	os.Remove(ServerSock)
}

func (s *Server) watch() {
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

func (s *Server) cachedExec(command string) (string, error) {
	if result, ok := s.resultCache[command]; ok {
		return result, nil
	}

	result, err := s.exec(command)
	if err != nil {
		return "", err
	}
	s.resultCache[command] = result

	return result, nil
}

func (s *Server) exec(command string) (string, error) {
	cmd := exec.Command("ghq", "list")

	result, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
