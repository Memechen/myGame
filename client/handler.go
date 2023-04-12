package client

import "github.com/Memechen/myGame/network"

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

func (c *Client) Login(param *InputParam) {

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
