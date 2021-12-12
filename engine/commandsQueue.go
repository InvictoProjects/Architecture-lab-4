package engine

import (
	"github.com/InvictoProjects/Architecture-lab-4/commands"
	"sync"
)

type commandsQueue struct {
	sync.Mutex
	commands []commands.Command
}

func (q *commandsQueue) push(cmd commands.Command) {
	q.Lock()
	defer q.Unlock()
	q.commands = append(q.commands, cmd)
}

func (q *commandsQueue) pull() commands.Command {
	q.Lock()
	defer q.Unlock()
	if len(q.commands) == 0 {
		return nil
	}
	res := q.commands[0]
	q.commands[0] = nil
	q.commands = q.commands[1:]
	return res
}

func (q *commandsQueue) isEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return len(q.commands) == 0
}
