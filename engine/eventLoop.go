package engine

import "github.com/InvictoProjects/Architecture-lab-4/commands"

type EventLoop struct {
	queue       *commandsQueue
	stopRequest bool
	stopConfirm chan struct{}
}

func (l *EventLoop) Post(cmd commands.Command) {
	l.queue.push(cmd)
}

func (l *EventLoop) Start() {
	l.queue = &commandsQueue{}
	l.stopConfirm = make(chan struct{})
	go func() {
		for {
			if l.stopRequest && l.queue.isEmpty() {
				break
			}
			cmd := l.queue.pull()
			if cmd != nil {
				cmd.Execute(l)
			}
		}
		l.stopConfirm <- struct{}{}
	}()
}

func (l *EventLoop) AwaitFinish() {
	l.stopRequest = true
	<-l.stopConfirm
}
