package network

import "github.com/Memechen/myGame/network/protocol/gen/messageId"

type Message struct {
	ID   messageId.MessageId
	Data []byte
}
