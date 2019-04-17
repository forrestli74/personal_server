package main

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
)

func stubCommand(x int) *tmp.Command {
	return &tmp.Command{
		Command: &tmp.Command_TickCommand{
			TickCommand: &tmp.TickCommand{
				RandomSeed: []byte{byte(x)},
			},
		},
	}
}

const N = 20

func TestHistoryCreateChan_IterateAllCommand(t *testing.T) {
	h := CreateHistory()
	ch := h.CreateChan(0)
	for i := 0; i < N/2; i++ {
		h.AppendCommand(stubCommand(i))
	}
	go func() {
		for i := N / 2; i < N; i++ {
			h.AppendCommand(stubCommand(i))
		}
	}()

	i := 0
	for i < N {
		commands := <-ch
		for _, command := range commands.Commands {
			if !proto.Equal(command, stubCommand(i)) {
				t.Fatalf("%d'th Command is not expected: %s", i, command)
			}
			i++
		}
	}

}

func TestHistory_CanCopy(t *testing.T) {
	h1 := CreateHistory()
	ch := h1.CreateChan(0)
	h2 := h1
	for i := 0; i < N; i++ {
		if i|1 == 0 {
			h1.AppendCommand(stubCommand(i))
		} else {
			h2.AppendCommand(stubCommand(i))
		}
	}

	i := 0
	for i < N {
		commands := <-ch
		for _, command := range commands.Commands {
			if !proto.Equal(command, stubCommand(i)) {
				t.Fatalf("%d'th Command is not expected: %s", i, command)
			}
			i++
		}
	}

}
