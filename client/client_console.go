package client

import (
	"bufio"
	"github.com/Memechen/myGame/log"
	"os"
	"strings"
)

type InputParam struct {
	Command string
	Param   []string
}

type ClientConsole struct {
	chInput chan *InputParam
}

func NewClientConsole() *ClientConsole {
	return &ClientConsole{}
}

func (c *ClientConsole) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Logger.ErrorF("input err, check your input and try again")
			continue
		}
		split := strings.Split(readString, " ")
		if len(split) == 0 {
			log.Logger.ErrorF("input err, check your input, your input is empty")
			continue
		}
		in := &InputParam{
			Command: split[0],
			Param:   []string{split[1], split[2]},
		}
		c.chInput <- in
	}
}
