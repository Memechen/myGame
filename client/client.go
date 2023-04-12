package client

import (
	"encoding/json"
	"github.com/Memechen/myGame/network"
)

type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[uint64]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParam
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:
				bytes, err := json.Marshal(input.Param)
				if err != nil {
					c.cli.ChMsg <- &network.Message{
						ID:   1,
						Data: bytes,
					}
				}
			}
		}
	}()
	go c.console.Run()
	go c.cli.Run()
}
