package main

import "github.com/Memechen/myGame/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}
