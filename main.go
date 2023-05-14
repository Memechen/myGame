package main

import (
	"github.com/Memechen/myGame/log"
	"github.com/Memechen/myGame/world"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	world.MM = world.NewMgrMgr()
	go world.MM.Run()
	log.Logger.InfoF("server start !!")
	WaitSignal(world.MM.OnSystemSignal)
}

func WaitSignal(fn func(sig os.Signal) bool) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE)
	for sig := range ch {
		if !fn(sig) {
			close(ch)
			break
		}
	}
}
