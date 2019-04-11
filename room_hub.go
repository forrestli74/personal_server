package main

import (
	"context"
	"fmt"
	"net/http"

	tmp "github.com/lijiaqigreat/personal_server/protobuf"
)

/*
RoomHub ...

*/
type RoomHub struct {
	roomByID map[string]*RoomServer
}

/*
Debug ...
*/
func (rh *RoomHub) Debug(c context.Context, request *tmp.DebugRequest) (*tmp.DebugResponse, error) {
	return nil, nil
}

/*
CreateRoom ...
*/
func (rh *RoomHub) CreateRoom(c context.Context, request *tmp.CreateRoomRequest) (*tmp.CreateRoomResponse, error) {
	if request.RoomId == "" {
		return nil, fmt.Errorf("RoomId cannot be empty")
	}
	if _, ok := rh.roomByID[request.RoomId]; ok {
		return nil, fmt.Errorf("Room already exists: %s", request.RoomId)
	}
	rh.roomByID[request.RoomId] = NewRoomServer(request.RoomSetting)
	return new(tmp.CreateRoomResponse), nil
}

/*
DeleteRoom ...
*/
func (rh *RoomHub) DeleteRoom(c context.Context, request *tmp.DeleteRoomRequest) (*tmp.DeleteRoomResponse, error) {
	rs, ok := rh.roomByID[request.RoomId]
	if !ok {
		return nil, fmt.Errorf("Room does not exist: %s", request.RoomId)
	}
	rs.Close()
	delete(rh.roomByID, request.RoomId)
	return new(tmp.DeleteRoomResponse), nil
}

/*
AddWriter ...
*/
func (rh *RoomHub) AddWriter(c context.Context, request *tmp.AddWriterRequest) (*tmp.AddWriterResponse, error) {
	roomServer, ok := rh.roomByID[request.RoomId]
	if !ok {
		return nil, fmt.Errorf("room not found")
	}
	return roomServer.AddWriter(c, request)
}

func (rh *RoomHub) Close() {
	for _, v := range rh.roomByID {
		v.Close()
	}
	rh.roomByID = make(map[string]*RoomServer)
}

/*
GetHandler ...
*/
func (rh *RoomHub) GetHandler() http.Handler {
	return roomHubHandler{rh: rh}
}

/*
NewRoomHub ...
*/
func NewRoomHub() *RoomHub {
	return &RoomHub{
		roomByID: make(map[string]*RoomServer),
	}
}

type roomHubHandler struct {
	rh *RoomHub
}

func (rhh roomHubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		http.Error(w, "Missing room_id", http.StatusBadRequest)
		return
	}
	if rs, ok := rhh.rh.roomByID[roomID]; ok {
		rs.GetHandler().ServeHTTP(w, r)
	} else {
		http.Error(w, "room_id not found", http.StatusBadRequest)
	}
}
