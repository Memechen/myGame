package network

import (
	"encoding/binary"
	"fmt"
	"github.com/Memechen/myGame/log"
	"net"
	"time"
)

type Client struct {
	Address   string
	packer    *NormalPacker
	ChMsg     chan *Message
	OnMessage func(message *ClientPacket)
}

func NewClient(address string) *Client {
	c := &Client{
		Address: address,
		packer:  NewNormalPacker(binary.BigEndian),
	}
	return c
}

func (c *Client) Run() {
	conn, err := net.Dial("tcp6", c.Address)
	if err != nil {
		log.Logger.ErrorF("net.Dial has err, c.Address: %s, err: %v", c.Address, err)
		return
	}
	go c.Write(conn)
	go c.Read(conn)
}

func (c *Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			c.send(conn, &Message{ID: 111, Data: []byte("hello, i am client")})
		}
	}
}

func (c Client) send(conn net.Conn, message *Message) {
	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Logger.ErrorF("conn.SetWriteDeadline has err, err: %v", err)
		return
	}
	bytes, err := c.packer.Pack(message)
	if err != nil {
		log.Logger.ErrorF("c.packer.Pack has err, err: %v", err)
		return
	}
	_, err = conn.Write(bytes)
	if err != nil {
		log.Logger.ErrorF("conn.Write(bytes) has err, err: %v", err)
		return
	}

}

func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)
		if _, ok := err.(net.Error); err != nil && ok {
			log.Logger.ErrorF("conn.SetWriteDeadline has err, err: %v", err)
			continue
		}
		c.OnMessage(&ClientPacket{
			Msg:  message,
			Conn: conn,
		})
		log.Logger.InfoF("resp message ID: ", fmt.Sprint(message.ID))
		log.Logger.InfoF("resp message Data: ", string(message.Data))
	}
}
