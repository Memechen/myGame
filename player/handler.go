package player

import (
	"github.com/Memechen/myGame/function"
	"github.com/Memechen/myGame/log"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/player"
	"github.com/golang/protobuf/proto"
)

type Handler func(packet *network.Message)

func (p *Player) AddFriend(packet *network.Message) {
	req := &player.CSAddFriend{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	if !function.CheckInNumberSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

func (p *Player) DelFriend(packet *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	p.FriendList = function.DelEleInSlice(req.UId, p.FriendList)
}

func (p *Player) ResolveChatMsg(packet *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	log.Logger.InfoF(req.Msg.Content)
}
