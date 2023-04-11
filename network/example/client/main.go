package main

import "github.com/Memechen/myGame/network"

func main() {
	client := network.NewClient(":8023")
	client.Run()
	select {}
}
