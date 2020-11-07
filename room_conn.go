package main

import (
	"net/http"
	"sync"

	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
)

/*
RoomConnState ...
*/
type RoomConnState struct {
}

/*
RoomConn ...
*/
type RoomConn struct {
	rs        *RoomServer
	id        string
	ch        <-chan []byte
	ws        *websocket.Conn
	closeOnce sync.Once
}

/*
Close ...
*/
func (rc *RoomConn) Close() {
	rc.closeOnce.Do(func() {
		rc.ws.Close()
		if rc.id != "" {
			rc.rs.appendRawCommand(&tmp.Command{
				Command: &tmp.Command_IdCommand{
					IdCommand: &tmp.IdCommand{
						OldId: rc.id,
					},
				},
			})
			delete(rc.rs.connectionByID, rc.id)
		}
		rc.ws = nil
	})
}

func CreateRoomConn(w http.ResponseWriter, r *http.Request, rs *RoomServer, id string, index int) (*RoomConn, error) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	rc := &RoomConn{rs: rs, id: id, ws: ws}

	if id != "" {
		rs.connectionByID[id] = rc
		rs.appendRawCommand(&tmp.Command{
			Command: &tmp.Command_IdCommand{
				IdCommand: &tmp.IdCommand{
					NewId: id,
				},
			},
		})
	}

	// handle command communication.
	commandChan := rc.rs.history.CreateChan(index)

	// write ws to history
	if rc.id != "" {
		go func() {
			defer rc.Close()
			for {
				_, message, err := ws.ReadMessage()
				if err != nil {
					break
				}
				rc.rs.appendRawCommand(&tmp.Command{
					Command: &tmp.Command_WriterCommand{
						WriterCommand: &tmp.WriterCommand{
							Id:      rc.id,
							Command: message,
						},
					},
				})
			}
		}()
	}

	// write history to ws
	go func() {
		defer rc.Close()
		for commands := range commandChan {
			bytes, _ := proto.Marshal(commands)
			err := ws.WriteMessage(websocket.BinaryMessage, bytes)
			if err != nil {
				break
			}
		}
	}()

	return rc, nil
}
