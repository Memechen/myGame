package world

import (
	"fmt"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/messageId"
	"github.com/Memechen/myGame/network/protocol/gen/player"
	logicPlayer "github.com/Memechen/myGame/player"
	"github.com/golang/protobuf/proto"
	"time"
)

func (mm *MgrMgr) CreateUser(message *network.SessionPacket) {
	msg := &player.CSCreateUser{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	fmt.Println("[CreateUser] server", msg)

	mm.SendMsg(message.Msg.ID, &player.SCCreateUser{}, message.Sess)
}

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	msg := &player.CSLogin{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}

	newPlayer := logicPlayer.NewPlayer()
	newPlayer.UId = uint64(time.Now().Unix())
	newPlayer.HandlerParamCh = message.Sess.WriteCh

	message.Sess.IsPlayerOnline = true
	mm.Pm.Add(newPlayer)
	newPlayer.Run()
}

func (mm MgrMgr) SendMsg(id messageId.MessageId, message proto.Message, session *network.Session) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	rsp := &network.Message{
		ID:   id,
		Data: bytes,
	}
	session.SendMsg(rsp)
}
