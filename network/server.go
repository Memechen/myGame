package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	address  string
	network  string
}

func NewServer(address, network string) *Server {
	return &Server{
		listener: nil,
		address:  address,
		network:  network,
	}
}

func (s *Server) Run() {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.address)
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.listener = tcpListener
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				continue
			}
			go func() {
				newSession := NewSession(conn)
				newSession.Run()
			}()
		}
	}()
}
