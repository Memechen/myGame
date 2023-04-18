package client

import (
	"fmt"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/player"
	"github.com/golang/protobuf/proto"
)

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

// 创建角色
func (c *Client) CreatePlayer(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 2 {
		return
	}
	msg := &player.CSCreateUser{
		Username: param.Param[0],
		Password: param.Param[1],
	}
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	c.cli.ChMsg <- &network.Message{
		ID:   uint64(id),
		Data: bytes,
	}
}

func (c Client) OnCreatePlayerRsp(packet *network.ClientPacket) {
	fmt.Println("恭喜你创建角色成功")
}

func (c *Client) Login(param *InputParam) {
	fmt.Println("[client handler Login] print")
	fmt.Println("[client handler Login] command：", param.Command)
	fmt.Println("[client handler Login] param：", param.Param)
}

func (c *Client) OnLoginResp(packet *network.ClientPacket) {

}

func (c *Client) AddFriend(param *InputParam) {

}

func (c *Client) OnAddFriendResp(packet *network.ClientPacket) {

}

func (c *Client) DelFriend(param *InputParam) {

}

func (c *Client) OnDelFriendResp(packet *network.ClientPacket) {

}

func (c *Client) SendChatMsg(param *InputParam) {

}

func (c *Client) OnSendChatMsgResp(packet *network.ClientPacket) {

}
