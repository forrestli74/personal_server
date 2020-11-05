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
	createdAt      time.Time
	closeCheckers  []RoomServerCloseChecker
}

/*
RoomServerCloseChecker ...
*/
type RoomServerCloseChecker interface {
	CheckClose(*RoomServer) bool
}

type maxDurationCloseChecker struct {
	closeTime time.Time
}

func (c *maxDurationCloseChecker) CheckClose(rs *RoomServer) bool {
	return time.Now().After(c.closeTime)
}

func newMaxDurationCloseChecker(rs *RoomServer) RoomServerCloseChecker {
	duration := rs.setting.GetEndOfLife().GetMaxDurationInNanoseconds()
	if duration == 0 {
		return nil
	}
	return &maxDurationCloseChecker{closeTime: time.Now().Add(time.Duration(duration) * time.Nanosecond)}
}

type closeWhenWriterDisconnectedChecker struct {
	connected bool
}

func (c *closeWhenWriterDisconnectedChecker) CheckClose(rs *RoomServer) bool {
	if len(rs.connectionByID) == 0 {
		return c.connected
	}
	c.connected = true
	return false
}

func newCloseWhenWriterDisconnectedChecker(rs *RoomServer) RoomServerCloseChecker {
	if rs.setting.GetEndOfLife().GetCloseWhenAllWriterDisconnected() {
		return &closeWhenWriterDisconnectedChecker{}
	}
	return nil
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
	rs.history.Close()
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

func (rs *RoomServer) checkIfNeedClose() (needClose bool) {
	return false
}

const maxNanoSecondLife = 24 * 3600 * 10e9

/*
NewRoomServer ...
*/
func NewRoomServer(setting *tmp.RoomSetting) (rs *RoomServer) {
	rs = &RoomServer{
		connectionByID: make(map[string]*RoomConn),
		history:        CreateHistory(),
		setting:        setting,
		closed:         false,
		closeCheckers:  []RoomServerCloseChecker{},
	}
	setupCloseCheckers(rs)

	period := setting.GetTick().GetFrequencyNanoseconds()
	duration := time.Duration(setting.GetEndOfLife().GetMaxDuration())
	if duration == 0 {
		duration = time.Duration(maxNanoSecondLife)
	}
	closeTime := time.Now().Add(duration)
	if period != 0 {
		ticker := time.NewTicker(time.Duration(period) * time.Nanosecond)
		go func() {
			randomBuffer := make([]byte, setting.GetTick().GetSize())
			for range ticker.C {
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

func setupCloseCheckers(rs *RoomServer) {
	var checker RoomServerCloseChecker

	checker = newMaxDurationCloseChecker(rs)
	if checker != nil {
		rs.closeCheckers = append(rs.closeCheckers, checker)
	}
	checker = newCloseWhenWriterDisconnectedChecker(rs)
	if checker != nil {
		rs.closeCheckers = append(rs.closeCheckers, checker)
	}
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
