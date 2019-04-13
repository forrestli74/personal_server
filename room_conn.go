package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
)

/*
RoomConnState ...
*/
type RoomConnState struct {
	ch <-chan []byte
	ws *websocket.Conn
}

/*
RoomConn ...
*/
type RoomConn struct {
	rs    *RoomServer
	id    string
	state RoomConnState
}

/*
Close ...
*/
func (rc *RoomConn) Close() {
	rc.state.ws.Close()
	if rc.id != "" {
		rc.rs.appendRawCommand(&tmp.Command{
			Command: &tmp.Command_IdCommand{
				IdCommand: &tmp.IdCommand{
					OldId: rc.id,
				},
			},
		})
	}
}

/*
Connect ...
*/
func (rc *RoomConn) Connect(w http.ResponseWriter, r *http.Request) error {
	if rc.state.ws != nil {
		http.Error(w, "Id not found", http.StatusBadRequest)
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	rc.state.ws = ws

	if rc.id != "" {
		rc.rs.appendRawCommand(&tmp.Command{
			Command: &tmp.Command_IdCommand{
				IdCommand: &tmp.IdCommand{
					NewId: rc.id,
				},
			},
		})
	}

	commandChan := rc.rs.history.CreateChan(0)

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
		for command := range commandChan {
			err := ws.WriteMessage(websocket.BinaryMessage, command)
			if err != nil {
				break
			}
		}
	}()

	return nil
}
