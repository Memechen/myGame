package world

import "github.com/Memechen/myGame/network/protocol/gen/messageId"

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageId.MessageId_CSLogin] = mm.UserLogin
}
