package player

import (
	"github.com/Memechen/myGame/define"
)

type Player struct {
	UId            uint64
	FriendList     []uint64
	HandlerParamCh chan define.HandlerParam
	handlers       map[string]Handler
}

func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: nil,
	}
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[handlerParam.HandlerKey]; ok {
				fn(handlerParam.Data)
			}
		}
	}
}
