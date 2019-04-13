package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	proto "github.com/golang/protobuf/proto"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
)

/*
RoomServer ...

*/
type RoomServer struct {
	setting        *tmp.RoomSetting
	connectionByID map[string]*RoomConn
	history        History
	closed         bool
}

/*
AddWriter ...
*/
func (rs *RoomServer) AddWriter(c context.Context, request *tmp.AddWriterRequest) (*tmp.AddWriterResponse, error) {
	for _, id := range request.ProposedIds {
		if _, ok := rs.connectionByID[id]; !ok {
			rs.connectionByID[id] = &RoomConn{id: id}
			rs.appendRawCommand(&tmp.Command{
				Command: &tmp.Command_IdCommand{
					IdCommand: &tmp.IdCommand{
						NewId: id,
					},
				},
			})
			return &tmp.AddWriterResponse{Id: id}, nil
		}
	}
	return &tmp.AddWriterResponse{}, nil
}

/*
GetHandler ...
*/
func (rs *RoomServer) GetHandler() http.Handler {
	return roomServerHandler{rs: rs}
}

/*
Close ...
*/
func (rs *RoomServer) Close() {
	rs.closed = true
}

/*
IsClosed ...
*/
func (rs *RoomServer) IsClosed() bool {
	return rs.closed
}

func (rs *RoomServer) appendRawCommand(command *tmp.Command) {
	rawCommand, _ := proto.Marshal(command)
	rs.history.AppendCommand(rawCommand)
}

/*
NewRoomServer ...
*/
func NewRoomServer(setting *tmp.RoomSetting) (rs *RoomServer) {
	rs = &RoomServer{
		connectionByID: make(map[string]*RoomConn),
		history:        CreateHistory(),
		setting:        setting,
		closed:         false,
	}
	period := setting.GetTickSetting().GetFrequencyMillis()
	if period != 0 {
		ticker := time.NewTicker(time.Duration(period) * time.Millisecond)
		go func() {
			randomBuffer := make([]byte, setting.GetTickSetting().GetSize())
			for range ticker.C {
				if rs.closed {
					break
				}
				rand.Read(randomBuffer)
				rs.appendRawCommand(&tmp.Command{
					Command: &tmp.Command_TickCommand{
						TickCommand: &tmp.TickCommand{
							RandomSeed: randomBuffer,
						},
					},
				})
			}
		}()
	}
	return
}

type roomServerHandler struct {
	rs *RoomServer
}

func (rsh roomServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	rc, ok := rsh.rs.connectionByID[id]
	if !ok {
		rc = &RoomConn{rs: rsh.rs, id: id}
		rsh.rs.connectionByID[id] = rc
	}

	rc.Connect(w, r)
}
