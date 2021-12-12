package tools

import (
	"github.com/InvictoProjects/Architecture-lab-4/commands"
	"strings"
)

func Parse(commandLine string) commands.Command {
	parts := strings.Fields(commandLine)
	if parts[0] == "print" {
		if len(parts) == 2 {
			return &commands.PrintCommand{Arg: parts[1]}
		} else {
			return &commands.PrintCommand{Arg: "SYNTAX ERROR: print command must have only one argument"}
		}
	} else if parts[0] == "cat" {
		if len(parts) == 3 {
			return &commands.CatCommand{Arg1: parts[1], Arg2: parts[2]}
		} else {
			return &commands.PrintCommand{Arg: "SYNTAX ERROR: cat command must have only two arguments"}
		}
	} else {
		return &commands.PrintCommand{Arg: "SYNTAX ERROR: could not parse " + commandLine}
	}
}
