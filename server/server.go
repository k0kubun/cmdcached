package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/howeyc/fsnotify"
)

const (
	ServerSock = "/tmp/cmdcached.sock"
	ConnType   = "unix"
	maxBuf     = 1024
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

	var buf [maxBuf]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}

	req := string(buf[:n])
	reqs := strings.SplitN(req, "\n", 2)
	if len(reqs) < 2 {
		fmt.Printf("Invalid request %s\n", req)
	}
	dir, cmd := reqs[0], reqs[1]

	result, err := s.cachedExec(dir, cmd)
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

func (s *Server) cachedExec(dir, command string) (string, error) {
	if result, ok := s.resultCache[command]; ok {
		return result, nil
	}

	result, err := s.exec(dir, command)
	if err != nil {
		return "", err
	}
	s.resultCache[command] = result

	return result, nil
}

func (s *Server) exec(dir, command string) (string, error) {
	err := os.Chdir(dir)
	if err != nil {
		return "", err
	}

	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)

	result, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
