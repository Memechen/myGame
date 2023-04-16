package player

import (
	"fmt"
	"github.com/Memechen/myGame/function"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/player"
	"github.com/golang/protobuf/proto"
)

type Handler func(packet *network.SessionPacket)

func (p *Player) AddFriend(packet *network.SessionPacket) {
	req := &player.CSAddFriend{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	if !function.CheckInNumberSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

func (p *Player) DelFriend(packet *network.SessionPacket) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = function.DelEleInSlice(req.UId, p.FriendList)
}

func (p *Player) ResolveChatMsg(packet *network.SessionPacket) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)
}
