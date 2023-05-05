package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sockapp/socket"
	"strings"
)

func say() error {
	conn, err := net.Dial("unix", socket.SockFile)
	if err != nil {
		return err
	}
	defer conn.Close()

	r := bufio.NewReader(os.Stdin)
	fmt.Printf("Say something: ")
	s, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	s = strings.TrimSpace(s)

	conn.Write([]byte(s))

	return nil
}

func main() {
	flag.Parse()

	sock := socket.New()

	var err error

	switch flag.Arg(0) {
	case "serve":
		sock.Listen()
	case "say":
		err = say()
	}

	if err != nil {
		panic(err)
	}
}
