package main

import (
	"testing"

	tmp "github.com/lijiaqigreat/personal_server/protobuf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestRoomServerSuite(t *testing.T) {
	suite.Run(t, new(RoomServerSuite))
}

type RoomServerSuite struct {
	suite.Suite
}

func (s *RoomServerSuite) TestSendsTick() {
	size := 2
	setting := &tmp.RoomSetting{
		Tick: &tmp.TickSetting{
			Size:                 uint32(size),
			FrequencyNanoseconds: 1000,
		},
	}
	rs := NewRoomServer(setting)
	defer rs.Close()
	ch := rs.history.CreateChan(0)
	commands := <-ch
	actual := commands.Commands[0]
	seed := actual.GetTickCommand().GetRandomSeed()
	assert.Equal(s.T(), len(seed), size)
}

func (s *RoomServerSuite) TestStopsAfterMaxDuration() {
	size := 2
	setting := &tmp.RoomSetting{
		Tick: &tmp.TickSetting{
			Size:                 uint32(size),
			FrequencyNanoseconds: 1e7,
		},
		EndOfLife: &tmp.EndOfLifeSetting{
			MaxDuration: 1,
		},
	}
	rs := NewRoomServer(setting)
	defer rs.Close()
	ch := rs.history.CreateChan(0)
	_, closed := <-ch
	assert.False(s.T(), closed)
}
