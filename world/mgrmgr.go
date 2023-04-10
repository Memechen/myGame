package world

import "github.com/Memechen/myGame/manager"

type MgrMgr struct {
	Pm manager.PlayerMgr
}

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
	return m
}

var MM *MgrMgr
