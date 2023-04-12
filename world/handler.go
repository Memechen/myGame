package world

import (
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/player"
)

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	newPlayer := player.NewPlayer()
	newPlayer.UId = 111
	newPlayer.HandlerParamCh = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	newPlayer.Run()
}
