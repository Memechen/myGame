package network

import (
	"encoding/binary"
	"fmt"
	"github.com/Memechen/myGame/chat"
	"net"
	"time"
)

type Client struct {
	Address   string
	packer    *NormalPacker
	ChMsg     chan chat.Msg
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
		fmt.Println(err)
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
		fmt.Println(err)
		return
	}
	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)
		if _, ok := err.(net.Error); err != nil && ok {
			fmt.Println(err)
			continue
		}
		c.OnMessage(&ClientPacket{
			Msg:  message,
			Conn: conn,
		})
		fmt.Println("resp message ID: ", fmt.Sprint(message.ID))
		fmt.Println("resp message Data: ", string(message.Data))
	}
}
