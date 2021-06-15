package main

import (
	"sync"

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
	Close()
}

type history2 struct {
	commands []*tmp.Command
	closed   bool
	/** always locked, but will be unlocked once a new locked lock is assigned. */
	rwLock *sync.RWMutex
	/** lock that protects operation on the history itself. */
	lock *sync.Mutex
}

/*
CreateHistory ...
*/
func CreateHistory() History {
	h := history2{
		rwLock: &sync.RWMutex{},
		lock:   &sync.Mutex{},
		closed: false,
	}
	h.rwLock.Lock()
	return &h
}

func (h *history2) AppendCommand(command *tmp.Command) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.closed {
		return
	}
	h.commands = append(h.commands, command)
	oldLock := h.rwLock
	newLock := &sync.RWMutex{}
	newLock.Lock()
	h.rwLock = newLock
	oldLock.Unlock()
}

func (h *history2) CreateChan(index int) <-chan *tmp.Commands {
	out := make(chan *tmp.Commands)

	go func() {
		for {
			mutex := h.rwLock
			length := len(h.commands)
			if index < length {
				out <- &tmp.Commands{
					Commands: h.commands[index:length],
				}
				index = length
			} else {
				mutex.RLock()
				if h.closed {
					close(out)
					break
				}
			}
		}
	}()
	return out
}

func (h *history2) Close() {
	h.lock.Lock()
	defer h.lock.Unlock()

	if !h.closed {
		h.closed = true
		h.rwLock.Unlock()
	}
}
