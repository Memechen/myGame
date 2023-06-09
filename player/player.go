package player

import (
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/messageId"
)

type Player struct {
	UId            uint64
	FriendList     []uint64
	HandlerParamCh chan *network.Message
	handlers       map[messageId.MessageId]Handler
	session        *network.Session
}

func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: nil,
		handlers:   make(map[messageId.MessageId]Handler),
	}
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[handlerParam.ID]; ok {
				fn(handlerParam)
			}
		}
	}
}
