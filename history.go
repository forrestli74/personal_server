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
CreateHistory2 ...
*/
func CreateHistory2() History {
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

type history2 struct {
	commands []*tmp.Command
	ins      map[chan struct{}]struct{}
	mutex    *sync.RWMutex
	out      chan<- *tmp.Command
}

/*
CreateHistory ...
*/
func CreateHistory() History {
	out := make(chan *tmp.Command)
	h := history2{
		out:   out,
		mutex: &sync.RWMutex{},
		ins:   map[chan struct{}]struct{}{},
	}
	h.mutex.Lock()
	return &h
}

func (h *history2) AppendCommand(command *tmp.Command) {
	h.commands = append(h.commands, command)
	m := h.mutex
	h.mutex = &sync.RWMutex{}
	h.mutex.Lock()
	m.Unlock()
}

func (h *history2) CreateChan(index int) <-chan *tmp.Commands {
	out := make(chan *tmp.Commands)

	go func() {
		for {
			length := len(h.commands)
			mutex := h.mutex
			if index < length {
				out <- &tmp.Commands{
					Commands: h.commands[index:length],
				}
				index = length
			} else {
				mutex.RLock()
			}
		}
	}()
	return out
}
