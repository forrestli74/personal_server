package main

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
	"github.com/stretchr/testify/suite"
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

func TestHistorySuite(t *testing.T) {
	suite.Run(t, new(HistorySuite))
}

type HistorySuite struct {
	suite.Suite
	history History
}

func (s *HistorySuite) SetupTest() {
	s.history = CreateHistory()
}

func (s *HistorySuite) TestIterateAllCommand() {
	ch := s.history.CreateChan(0)
	for i := 0; i < N/2; i++ {
		s.history.AppendCommand(stubCommand(i))
	}
	go func() {
		for i := N / 2; i < N; i++ {
			s.history.AppendCommand(stubCommand(i))
		}
	}()

	i := 0
	for i < N {
		commands := <-ch
		for _, command := range commands.Commands {
			if !proto.Equal(command, stubCommand(i)) {
				s.T().Fatalf("%d'th Command is not expected: %s", i, command)
			}
			i++
		}
	}

}

func (s *HistorySuite) TestCanCopy() {
	ch := s.history.CreateChan(0)
	cp := s.history
	for i := 0; i < N; i++ {
		if i|1 == 0 {
			s.history.AppendCommand(stubCommand(i))
		} else {
			cp.AppendCommand(stubCommand(i))
		}
	}

	i := 0
	for i < N {
		commands := <-ch
		for _, command := range commands.Commands {
			if !proto.Equal(command, stubCommand(i)) {
				s.T().Fatalf("%d'th Command is not expected: %s", i, command)
			}
			i++
		}
	}
}
