package connection

import (
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

type WebSocketConnection struct {
	Conn     *websocket.Conn
	CtxClose context.CancelFunc
	mutex    sync.Mutex
}

func (ws *WebSocketConnection) GetID() string {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	return ws.Conn.RemoteAddr().String()
}

func (ws *WebSocketConnection) Write(msgType int, msg []byte) error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	err := ws.Conn.WriteMessage(msgType, msg)
	return err
}

func (ws *WebSocketConnection) Close() {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	if ws.CtxClose != nil {
		ws.CtxClose()
	}
}
