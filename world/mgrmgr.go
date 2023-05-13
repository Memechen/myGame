package world

import (
	"fmt"
	"github.com/Memechen/myGame/manager"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/messageId"
	"os"
	"syscall"
)

type MgrMgr struct {
	Pm              *manager.PlayerMgr
	Server          *network.Server
	Handlers        map[messageId.MessageId]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{Pm: &manager.PlayerMgr{}}
	m.Server = network.NewServer(":8023")
	m.Server.OnSessionPacket = m.OnSessionPacket
	return m
}

var MM *MgrMgr

func (mm *MgrMgr) Run() {
	go mm.Server.Run()
	go mm.Pm.Run()
}

func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := mm.Handlers[packet.Msg.ID]; ok {
		handler(packet)
	}
	if p := mm.Pm.GetPlayer(packet.Sess.UId); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

func (mm MgrMgr) OnSystemSignal(signal os.Signal) bool {
	fmt.Printf("[Client] 收到信号 %v \n", signal)
	tag := true
	switch signal {
	case syscall.SIGHUP:
		// todo 挂起的时候做点事情
	case syscall.SIGPIPE:

	default:
		fmt.Println("[Client] 收到信号准备退出...")
		tag = false
	}
	return tag
}
