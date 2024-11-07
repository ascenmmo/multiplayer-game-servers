package connection

import (
	"context"
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	Conn     *websocket.Conn
	CtxClose context.CancelFunc
}

func (u *WebSocketConnection) GetID() string {
	return u.Conn.RemoteAddr().String()
}

func (u *WebSocketConnection) Write(msg []byte) error {
	err := u.Conn.WriteMessage(websocket.BinaryMessage, msg)

	return err
}

func (u *WebSocketConnection) Close() {
	if u.CtxClose != nil {
		u.CtxClose()
	}
}
