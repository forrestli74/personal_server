package main

import (
	"testing"

	tmp "github.com/lijiaqigreat/personal_server/protobuf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestRoomHubSuite(t *testing.T) {
	suite.Run(t, new(RoomHubSuite))
}

type RoomHubSuite struct {
	suite.Suite
	rh *RoomHub
}

func (s *RoomHubSuite) SetupTest() {
	s.rh = NewRoomHub()
}
func (s *RoomHubSuite) TearDownTest() {
	s.rh.Close()
}

func (s *RoomHubSuite) TestCreateRoomWorks() {
	id := "id"
	setting := &tmp.RoomSetting{
		Tick: &tmp.TickSetting{
			Size:                 2,
			FrequencyNanoseconds: 1e7,
		},
	}
	response, err := s.rh.CreateRoom(nil, &tmp.CreateRoomRequest{
		RoomId:      id,
		RoomSetting: setting,
	})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), *response, tmp.CreateRoomResponse{})

	//assert.Equal(s.T(), s.rh.roomByID[id].setting, setting)
	assert.Equal(s.T(), s.rh.roomByID[id].server.setting, setting)
}

func (s *RoomHubSuite) TestCreateRoomWithEmptyRoomIdFails() {
	setting := &tmp.RoomSetting{
		Tick: &tmp.TickSetting{
			Size:                 2,
			FrequencyNanoseconds: 1e7,
		},
	}
	_, err := s.rh.CreateRoom(nil, &tmp.CreateRoomRequest{
		RoomSetting: setting,
	})
	assert.NotNil(s.T(), err)
}

func (s *RoomHubSuite) TestCreateRoomWithExistingRoomIdFails() {
	id := "id"
	s.rh.CreateRoom(nil, &tmp.CreateRoomRequest{
		RoomId: id,
	})

	setting := &tmp.RoomSetting{
		Tick: &tmp.TickSetting{
			Size:                 2,
			FrequencyNanoseconds: 1e7,
		},
	}
	_, err := s.rh.CreateRoom(nil, &tmp.CreateRoomRequest{
		RoomId:      id,
		RoomSetting: setting,
	})
	assert.NotNil(s.T(), err)
	//assert.Nil(s.T(), s.rh.roomByID[id].setting)
	assert.Nil(s.T(), s.rh.roomByID[id].server.setting)
}

func (s *RoomHubSuite) TestDeleteRoomWorks() {
	id := "id"
	s.rh.CreateRoom(nil, &tmp.CreateRoomRequest{
		RoomId: id,
	})
	rs := s.rh.roomByID[id].server

	_, err := s.rh.DeleteRoom(nil, &tmp.DeleteRoomRequest{
		RoomId: id,
	})
	assert.Nil(s.T(), err)
	assert.Nil(s.T(), s.rh.roomByID[id])
	assert.True(s.T(), rs.IsClosed())
}

func (s *RoomHubSuite) TestListRoomWorks() {
	id := "id"
	setting := &tmp.RoomSetting{
		Tick: &tmp.TickSetting{
			Size: 1,
		},
	}
	description := "description"
	s.rh.CreateRoom(nil, &tmp.CreateRoomRequest{
		RoomId:           id,
		ShortDescription: description,
		RoomSetting:      setting,
	})

	response, err := s.rh.ListRoom(nil, nil)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), response, &tmp.ListRoomResponse{
		Rooms: []*tmp.RoomSummary{&tmp.RoomSummary{Id: id, Setting: setting, ShortDescription: description}},
	})
}
