package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	rs.history.AppendCommand(command)
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
	period := setting.GetTick().GetFrequencyMillis()
	duration := time.Duration(setting.GetEndOfLife().GetMaxDurationInSeconds() * 1000)
	if duration == 0 {
		duration = time.Duration(1000000000)
	}
	closeTime := time.Now().Add(duration)
	if period != 0 {
		ticker := time.NewTicker(time.Duration(period) * time.Millisecond)
		go func() {
			randomBuffer := make([]byte, setting.GetTick().GetSize())
			for tickTime := range ticker.C {
				if tickTime.After(closeTime) {
					rs.Close()
					break
				}
				if rs.IsClosed() {
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
	query := r.URL.Query()
	id := query.Get("id")
	index, _ := strconv.Atoi(query.Get("index"))
	rc, ok := rsh.rs.connectionByID[id]
	if id == "" {
		rc = &RoomConn{rs: rsh.rs}
	} else if !ok {
		rc = &RoomConn{rs: rsh.rs, id: id}
		rsh.rs.connectionByID[id] = rc
	}

	rc.Connect(w, r, index)
}
