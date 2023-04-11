package network

import (
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn   net.Conn
	packer *NormalPacker
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
	}
}

func (s *Session) Run() {

}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		return
	}
	message, err := s.packer.UnPack(s.conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server receive message: ", string(message.Data))
}

func (s *Session) Write(data []byte) {
	s.conn.SetWriteDeadline(time.Now().Add(time.Second))
}
