package connection

import (
	"encoding/json"
	"github.com/ascenmmo/websocket-server/pkg/restconnection/types"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
)

type NotifyServers interface {
	NotifyServers(ids []uuid.UUID, req types.Request) error
	AddServer(ID uuid.UUID, token string, addr string) error
}

type notifier struct {
	servers []*server
}

func NewNotifierServers() NotifyServers {
	return &notifier{}
}

type server struct {
	ID         uuid.UUID
	Addr       string
	Connection *websocket.Conn
}

func (n *notifier) NotifyServers(ids []uuid.UUID, req types.Request) error {
	if len(n.servers) == 0 {
		return nil
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	for _, id := range ids {
		for i, server := range n.servers {
			if server.ID == id {
				err = n.servers[i].Connection.WriteMessage(websocket.BinaryMessage, marshal)
				if err != nil {
					err = n.servers[i].Connect(req.Token)
					if err != nil {
						n.RemoveNotifyServer(id)
						return err
					}
					return n.servers[i].Connection.WriteMessage(websocket.BinaryMessage, marshal)
				}
			}
		}
	}
	return nil
}

func (n *notifier) AddServer(ID uuid.UUID, token string, addr string) error {
	newServer := &server{
		ID:   ID,
		Addr: addr,
	}
	err := newServer.Connect(token)
	if err != nil {
		return err
	}
	for i, s := range n.servers {
		if s.ID == ID {
			n.servers[i] = newServer
			return nil
		}
	}
	n.servers = append(n.servers, newServer)
	return nil
}

func (n *notifier) RemoveNotifyServer(id uuid.UUID) {
	for i, s := range n.servers {
		if s.ID == id {
			_ = n.servers[i].Connection.Close()
			n.servers = append(n.servers[:i], n.servers[i+1:]...)
		}
	}
}

func (s *server) Connect(token string) error {
	url := s.Addr + "/api/ws/connect"

	headers := http.Header{}
	headers.Add("token", token)

	conn, _, err := websocket.DefaultDialer.Dial(url, headers)
	if err != nil {
		return err
	}
	s.Connection = conn

	return nil
}
