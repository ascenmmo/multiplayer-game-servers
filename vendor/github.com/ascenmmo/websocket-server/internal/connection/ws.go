package connection

import "github.com/gorilla/websocket"

type WebSocketConnection struct {
	Conn *websocket.Conn
}

func (u *WebSocketConnection) GetID() string {
	return u.Conn.RemoteAddr().String()
}

func (u *WebSocketConnection) Write(msg []byte) error {
	err := u.Conn.WriteMessage(websocket.BinaryMessage, msg)

	return err
}
