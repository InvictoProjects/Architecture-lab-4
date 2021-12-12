package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/InvictoProjects/Architecture-lab-4/engine"
	"github.com/InvictoProjects/Architecture-lab-4/tools"
)

const inputFile = "./test/commands.txt"

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := strings.TrimSpace(scanner.Text())
			if commandLine == "" {
				continue
			}
			cmd := tools.Parse(commandLine)
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
