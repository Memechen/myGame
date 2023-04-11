package player

import (
	"fmt"
	"github.com/Memechen/myGame/chat"
	"github.com/Memechen/myGame/function"
)

type Handler func(interface{})

func (p *Player) AddFriend(data interface{}) {
	fId := data.(uint64)
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFriend(data interface{}) {
	fId := data.(uint64)
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

func (p *Player) ResolveChatMsg(data interface{}) {
	chatMsg := data.(chat.Msg)
	fmt.Println(chatMsg)
}
