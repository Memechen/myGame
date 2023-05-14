package network

import (
	"encoding/binary"
	"github.com/Memechen/myGame/log"
	"net"
	"time"
)

type Session struct {
	UId            uint64
	conn           net.Conn
	IsClose        bool
	packer         *NormalPacker
	WriteCh        chan *Message
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		WriteCh: make(chan *Message, 1),
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
			log.Logger.ErrorF("s.packer.UnPack has err, err: %v", err)
			continue
		}
		log.Logger.InfoF("server receive message: ", string(message.Data))
		s.MessageHandler(&SessionPacket{
			Msg:  message,
			Sess: s,
		})
		s.WriteCh <- &Message{
			ID:   999,
			Data: []byte("hi client, i am hulk"),
		}

	}
}

func (s *Session) Write() {
	for {
		select {
		case msg := <-s.WriteCh:
			s.send(msg)
		}
	}
}

func (s *Session) send(message *Message) {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
	if err != nil {
		log.Logger.ErrorF("s.conn.SetWriteDeadline has err, err: %v", err)
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

func (s *Session) SendMsg(msg *Message) {
	s.WriteCh <- msg
}
