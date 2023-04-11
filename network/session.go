package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn    net.Conn
	packer  *NormalPacker
	chWrite chan *Message
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		chWrite: make(chan *Message, 1),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Read() {
	for {
		err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			return
		}
		message, err := s.packer.UnPack(s.conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("server receive message: ", string(message.Data))
		s.chWrite <- &Message{
			Id:   999,
			Data: []byte("hi client, i am hulk"),
		}
	}
}

func (s *Session) Write() {

	for {
		select {
		case msg := <-s.chWrite:
			s.send(msg)
		}
	}
}

func (s *Session) send(message *Message) {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := s.packer.Pack(message)
	if err != nil {
		return
	}
	_, err = s.conn.Write(bytes)
	if err != nil {
		return
	}
}