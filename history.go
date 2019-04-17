package main

import (
	"sync"

	proto "github.com/golang/protobuf/proto"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
)

/*
RawCommand ...
*/
type RawCommand []byte

type history struct {
	commands []RawCommand
	mutex    *sync.RWMutex
}

/*
History ...
*/
type History interface {
	AppendCommand(command *tmp.Command)
	CreateChan(index int) <-chan *tmp.Commands
}

/*
CreateHistory ...
*/
func CreateHistory() History {
	h := history{
		mutex: &sync.RWMutex{},
	}
	h.mutex.Lock()
	return &h
}

func (h *history) AppendCommand(command *tmp.Command) {
	raw, _ := proto.Marshal(command)
	h.commands = append(h.commands, raw)
	h.mutex.Unlock()
	h.mutex = &sync.RWMutex{}
	h.mutex.Lock()
}

func (h *history) CreateChan(index int) <-chan *tmp.Commands {
	out := make(chan *tmp.Commands)
	go func() {
		for len(h.commands) <= index {
			h.mutex.RLock()
		}
		for {
			commands := make([]*tmp.Command, 0, len(h.commands)-index)
			for index < len(h.commands) {
				command := new(tmp.Command)
				proto.Unmarshal(h.commands[index], command)
				commands = append(commands, command)

				index++
			}
			out <- &tmp.Commands{
				Commands: commands,
			}
			h.mutex.RLock()
			if index == len(h.commands) {
				close(out)
				break
			}
		}
	}()
	return out
}
