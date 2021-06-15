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
        tmp.UnsafeRoomServiceServer 
	roomByID map[string]*room
}

type room struct {
	server           *RoomServer
	shortDescription string
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
	rh.roomByID[request.RoomId] = &room{
		server:           NewRoomServer(request.RoomSetting),
		shortDescription: request.ShortDescription,
	}
	return new(tmp.CreateRoomResponse), nil
}

/*
DeleteRoom ...
*/
func (rh *RoomHub) DeleteRoom(c context.Context, request *tmp.DeleteRoomRequest) (*tmp.DeleteRoomResponse, error) {
	room, ok := rh.roomByID[request.RoomId]
	if !ok {
		return nil, fmt.Errorf("Room does not exist: %s", request.RoomId)
	}
	room.server.Close()
	delete(rh.roomByID, request.RoomId)
	return new(tmp.DeleteRoomResponse), nil
}

/*
ListRoom ...
*/
func (rh *RoomHub) ListRoom(c context.Context, request *tmp.ListRoomRequest) (*tmp.ListRoomResponse, error) {
	var rooms []*tmp.RoomSummary
	for k, v := range rh.roomByID {
		rooms = append(rooms, &tmp.RoomSummary{
			Id:               k,
			Setting:          v.server.setting,
			ShortDescription: v.shortDescription,
		})
	}
	return &tmp.ListRoomResponse{
		Rooms: rooms,
	}, nil
}
func (rh *RoomHub) mustEmbedUnimplementedRoomServiceServer() {}

/*
Close ...
*/
func (rh *RoomHub) Close() {
	for _, v := range rh.roomByID {
		v.server.Close()
	}
	rh.roomByID = make(map[string]*room)
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
		roomByID: make(map[string]*room),
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
	if room, ok := rhh.rh.roomByID[roomID]; ok {
		room.server.GetHandler().ServeHTTP(w, r)
	} else {
		http.Error(w, "room_id not found", http.StatusBadRequest)
	}
}
