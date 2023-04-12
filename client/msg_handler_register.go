package client

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[111] = c.OnLoginResp
	c.messageHandlers[222] = c.OnAddFriendResp
	c.messageHandlers[333] = c.OnDelFriendResp
	c.messageHandlers[444] = c.OnSendChatMsgResp

}
