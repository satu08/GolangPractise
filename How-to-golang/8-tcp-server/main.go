package main

import (
	"fmt"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	ipaddress string
	ln        net.Listener
	quitchan  chan struct{}
	msgchan   chan Message
}

func NewServer(ipaddress string) *Server {
	return &Server{
		ipaddress: ipaddress,
		quitchan:  make(chan struct{}),
		msgchan:   make(chan Message, 1024),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ipaddress)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.acceptLoop()
	<-s.quitchan
	close(s.msgchan)
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}
		fmt.Println("accepting connection from ", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading from connection:", err)
		}
		s.msgchan <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
		conn.Write([]byte("thank you for your msg!"))
	}
}

func main() {
	server := NewServer(":8080")
	go func() {
		for msg := range server.msgchan {
			fmt.Printf("received msg from %s:%s\n", msg.from, string(msg.payload))
		}
	}()
	server.Start()
}
