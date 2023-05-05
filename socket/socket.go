package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	SockFile = "/tmp/app.sock"
)

type Socket struct {
	addr string
}

func (s *Socket) Quit() {
	os.Remove(SockFile)
	os.Exit(0)
}

func (sock *Socket) Listen() error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	go func() {
		select {
		case <-sigCh:
			sock.Quit()
		}
	}()

	ln, err := net.Listen("unix", sock.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		s := bufio.NewScanner(conn)
		for s.Scan() {
			t := s.Text()

			switch t {
			case "q":
				sock.Quit()

			default:
				fmt.Printf("Client said: '%s'\n", t)
			}
		}
	}
}

func New() *Socket {
	return &Socket{
		addr: SockFile,
	}
}
