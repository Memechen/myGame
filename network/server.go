package network

import (
	"github.com/Memechen/myGame/log"
	"net"
)

type Server struct {
	listener        net.Listener
	address         string
	OnSessionPacket func(packet *SessionPacket)
}

func NewServer(address string) *Server {
	return &Server{
		listener: nil,
		address:  address,
	}
}

func (s *Server) Run() {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.address)
	if err != nil {
		log.Logger.ErrorF("net.ResolveTCPAddr: %s has err, err: %v", s.address, err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		log.Logger.ErrorF("net.ListenTCP: %s has err, err: %v", resolveTCPAddr, err)
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
