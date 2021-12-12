package commands

import "fmt"

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type PrintCommand struct {
	Arg string
}

func (pc *PrintCommand) Execute(handler Handler) {
	fmt.Println(pc.Arg)
}
