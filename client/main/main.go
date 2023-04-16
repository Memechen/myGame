package main

import "github.com/Memechen/myGame/client"

func main() {
	c := client.NewClient()
	c.InputHandlerRegister()
	c.Run()
	select {}
}
