package client

import (
	"fmt"
	"github.com/Memechen/myGame/network"
	"github.com/Memechen/myGame/network/protocol/gen/messageId"
	"os"
	"syscall"
)

type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[messageId.MessageId]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParam
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewClientConsole(),
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	return c
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:
				fmt.Printf("[client run] cmd:%s, param:%v <<<\t \n", input.Command, input.Command)
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}

			}
		}
	}()
	go c.console.Run()
	go c.cli.Run()
}

func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[packet.Msg.ID]; ok {
		handler(packet)
	}
}

func (c *Client) OnSystemSignal(signal os.Signal) bool {
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
