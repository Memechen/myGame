package client

import "github.com/Memechen/myGame/network/protocol/gen/messageId"

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[messageId.MessageId_SCCreatePlayer] = c.OnCreatePlayerRsp
	c.messageHandlers[messageId.MessageId_SCLogin] = c.OnLoginResp
	c.messageHandlers[messageId.MessageId_SCAddFriend] = c.OnAddFriendResp
	c.messageHandlers[messageId.MessageId_SCDelFriend] = c.OnDelFriendResp
	c.messageHandlers[messageId.MessageId_SCSendChatMsg] = c.OnSendChatMsgResp
}
