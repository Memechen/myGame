package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	Order binary.ByteOrder
}

func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{
		Order: order,
	}
}

// Pack |data 长度|id|data|
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buffer := make([]byte, 8+8+len(message.Data))
	p.Order.PutUint64(buffer[:8], uint64(len(buffer)))
	p.Order.PutUint64(buffer[8:16], message.Id)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

func (p *NormalPacker) UnPack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	totalLen := p.Order.Uint64(buffer[:8])
	Id := p.Order.Uint64(buffer[8:])
	dataSize := totalLen - 8 - 8
	dataBuffer := make([]byte, dataSize)
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	m := &Message{
		Id:   Id,
		Data: dataBuffer,
	}
	return m, nil
}