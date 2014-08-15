package main

import (
	"bufio"
	"github.com/lisijie/go-util"
	"log"
	"net"
)

type EchoServer struct {
	listener net.Listener
}

func (s *EchoServer) Handle(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	wd := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		wd.Write(line)
		wd.WriteByte('\n')
		wd.Flush()
	}
}

func (s *EchoServer) Start() {
	util.TCPServer(s.listener, s)
}

func main() {
	listener, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal(err)
	}

	server := &EchoServer{listener: listener}
	server.Start()
}
