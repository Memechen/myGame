package client

import (
	"github.com/Memechen/myGame/log"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/player"
	"github.com/golang/protobuf/proto"
	"strconv"
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

	c.Transport(id, msg)
}

func (c Client) OnCreatePlayerRsp(packet *network.ClientPacket) {
	log.Logger.InfoF("恭喜你创建角色成功")
}

func (c *Client) Login(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &player.CSLogin{
		Username: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)
}

func (c *Client) OnLoginResp(packet *network.ClientPacket) {
	resp := &player.SCLogin{}

	err := proto.Unmarshal(packet.Msg.Data, resp)
	if err != nil {
		return
	}

	log.Logger.InfoF("登录成功")
}

func (c *Client) AddFriend(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSAddFriend{
		UId: parseUint,
	}

	c.Transport(id, msg)
}

func (c *Client) OnAddFriendResp(packet *network.ClientPacket) {

}

func (c *Client) DelFriend(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSDelFriend{
		UId: parseUint,
	}

	c.Transport(id, msg)
}

func (c *Client) OnDelFriendResp(packet *network.ClientPacket) {
	log.Logger.InfoF("you have del friend success")
}

func (c *Client) SendChatMsg(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	parseUint32, err := strconv.ParseUint(param.Param[2], 10, 32)
	if err != nil {
		return
	}

	msg := &player.CSSendChatMsg{
		UId: parseUint,
		Msg: &player.ChatMessage{
			Content: param.Param[1],
			Extra:   nil,
		},
		Category: int32(parseUint32),
	}

	c.Transport(id, msg)
}

func (c *Client) OnSendChatMsgResp(packet *network.ClientPacket) {
	log.Logger.InfoF("send chat msg success")
}
