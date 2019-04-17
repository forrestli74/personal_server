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
		TickSetting: &tmp.TickSetting{
			Size:            uint32(size),
			FrequencyMillis: 1,
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
