package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/InvictoProjects/Architecture-lab-4/engine"
	"github.com/InvictoProjects/Architecture-lab-4/tools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide file with commands as argument.")
	} else if input, err := os.Open(os.Args[1]); err == nil {
		eventLoop := new(engine.EventLoop)
		eventLoop.Start()
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
		eventLoop.AwaitFinish()
	} else {
		fmt.Println("File does not exist.")
	}

}
