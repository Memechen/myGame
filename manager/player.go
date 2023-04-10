package manager

import "github.com/Memechen/myGame/player"

// PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]player.Player
	addPh   chan player.Player
}

// Add 添加方法 不会并发报错吗？
func (pm *PlayerMgr) Add(p player.Player) {
	pm.players[p.UId] = p
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPh:
			pm.Add(p)
		}
	}
}
