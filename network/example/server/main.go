package main

import "github.com/Memechen/myGame/network"

func main() {
	server := network.NewServer(":8023", "tcp6")
	server.Run()
	select {}
}
