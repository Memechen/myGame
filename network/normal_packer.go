package network

import (
	"encoding/binary"
	"github.com/Memechen/myGame/log"
	"github.com/Memechen/myGame/network/protocol/gen/messageId"
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
	p.Order.PutUint64(buffer[8:16], uint64(message.ID))
	copy(buffer[16:], message.Data)
	return buffer, nil
}

func (p *NormalPacker) UnPack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Logger.ErrorF("reader.(*net.TCPConn).SetReadDeadline has err, err: %v", err)
		return nil, err
	}
	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		log.Logger.ErrorF("io.ReadFull(reader, buffer) has err, err: %v", err)
		return nil, err
	}
	totalLen := p.Order.Uint64(buffer[:8])
	Id := p.Order.Uint64(buffer[8:])
	dataSize := totalLen - 8 - 8
	dataBuffer := make([]byte, dataSize)
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		log.Logger.ErrorF("io.ReadFull(reader, dataBuffer) has err, err: %v", err)
		return nil, err
	}
	m := &Message{
		ID:   messageId.MessageId(Id),
		Data: dataBuffer,
	}
	return m, nil
}
