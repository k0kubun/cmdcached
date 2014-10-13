package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const (
	ServerSock = "/tmp/cmdcached.sock"
	ConnType   = "unix"
	maxBuf     = 1024
)

type Server struct {
	config      *Config
	cmdCache    map[string]string
	dirCmdCache map[string]string
	listener    *net.UnixListener
	subscriber  *Subscriber
}

func NewServer() *Server {
	s := new(Server)
	s.config = NewConfig()
	s.subscriber = NewSubscriber()
	s.cmdCache = make(map[string]string)
	s.dirCmdCache = make(map[string]string)

	os.Remove(ServerSock) // avoid "address already in use"
	l, err := net.ListenUnix(
		ConnType,
		&net.UnixAddr{ServerSock, ConnType},
	)
	if err != nil {
		log.Println(err)
	}
	s.listener = l

	go s.cacheSweeper()

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

func (s *Server) cachedExec(dir, cmd string) (string, error) {
	if s.config.DirIgnorable(cmd) {
		if result, ok := s.cmdCache[cmd]; ok {
			return result, nil
		}
	} else {
		if result, ok := s.dirCmdCache[dir+"\n"+cmd]; ok {
			return result, nil
		}
	}

	result, err := s.exec(dir, cmd)
	if err != nil {
		return "", err
	}

	if s.config.DirIgnorable(cmd) {
		s.cmdCache[cmd] = result
	} else {
		s.dirCmdCache[dir+"\n"+cmd] = result
	}

	return result, nil
}

func (s *Server) exec(dir, cmd string) (string, error) {
	if !s.config.DirIgnorable(cmd) {
		err := os.Chdir(dir)
		if err != nil {
			return "", err
		}
	}

	args := strings.Split(cmd, " ")
	c := exec.Command(args[0], args[1:]...)

	result, err := c.Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (s *Server) cacheSweeper() {
	for {
		select {
		case <-s.subscriber.Events:
			// purge
		}
	}
}
