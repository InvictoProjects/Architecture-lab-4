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

type CatCommand struct {
	Arg1, Arg2 string
}

func (cat *CatCommand) Execute(handler Handler) {
	res := cat.Arg1 + cat.Arg2
	handler.Post(&PrintCommand{Arg: res})
}
