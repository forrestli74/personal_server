package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
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
	observers      []chan<- *tmp.Command
	closeObservers []chan<- string
	closeOnce      sync.Once
}

/*
AddObserver ...
*/
func (rs *RoomServer) AddObserver() <-chan *tmp.Command {
	ch := make(chan *tmp.Command)
	rs.observers = append(rs.observers, ch)
	return ch
}

/*
AddCloseObserver ...
*/
func (rs *RoomServer) AddCloseObserver() <-chan string {
	ch := make(chan string)
	rs.closeObservers = append(rs.closeObservers, ch)
	return ch
}

func (rs *RoomServer) setupMaxDurationCheck() {
	duration := rs.setting.GetEndOfLife().GetMaxDuration()
	if duration == 0 {
		return
	}
	timer := time.NewTimer(time.Duration(duration))
	closed := rs.AddCloseObserver()
	select {
	case <-timer.C:
		rs.Close()
		<-closed
	case <-closed:
		timer.Stop()
		rs.Close()
	}
}

func (rs *RoomServer) setupCloseWhenEmptyCheck() {
	if !rs.setting.GetEndOfLife().GetCloseWhenAllWriterDisconnected() {
		return
	}
	connected := false
	ch := rs.AddObserver()
	go func() {
		for range ch {
			if len(rs.connectionByID) != 0 {
				connected = true
				continue
			}
			if !connected {
				continue
			}
			break
		}
		rs.Close()
	}()
}

func (rs *RoomServer) setupTicker() {
	period := rs.setting.GetTick().GetFrequencyNanoseconds()
	if period == 0 {
		return
	}
	ticker := time.NewTicker(time.Duration(period) * time.Nanosecond)
	closed := rs.AddCloseObserver()
	go func() {
		<-closed
		ticker.Stop()
	}()
	size := rs.setting.GetTick().GetSize()
	alwaysActive := rs.setting.GetTick().GetAlwaysActive()

	go func() {
		randomBuffer := make([]byte, size)
		for range ticker.C {
			if rs.IsClosed() {
				break
			}
			if alwaysActive || (len(rs.connectionByID) > 0) {
				rand.Read(randomBuffer)
				rs.appendRawCommand(&tmp.Command{
					Command: &tmp.Command_TickCommand{
						TickCommand: &tmp.TickCommand{
							RandomSeed: randomBuffer,
						},
					},
				})
			}
		}
	}()
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
	rs.closeOnce.Do(func() {
		rs.closed = true
		for _, ch := range rs.observers {
			close(ch)
		}
		for _, ch := range rs.closeObservers {
			close(ch)
		}

		rs.history.Close()
	})
}

/*
IsClosed ...
*/
func (rs *RoomServer) IsClosed() bool {
	return rs.closed
}

func (rs *RoomServer) appendRawCommand(command *tmp.Command) {
	rs.history.AppendCommand(command)
	for _, ch := range rs.observers {
		ch <- command
	}
}

/*
NewRoomServer ...
*/
func NewRoomServer(setting *tmp.RoomSetting) (rs *RoomServer) {
	rs = &RoomServer{
		connectionByID: make(map[string]*RoomConn),
		history:        CreateHistory(),
		setting:        setting,
	}
	// setupCloseCheckers(rs)

	rs.setupCloseWhenEmptyCheck()
	rs.setupMaxDurationCheck()
	rs.setupTicker()
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
	_, ok := rsh.rs.connectionByID[id]
	if ok {
		http.Error(w, "Already connected by someone else", http.StatusBadRequest)
	}

	CreateRoomConn(w, r, rsh.rs, id, index)
}
