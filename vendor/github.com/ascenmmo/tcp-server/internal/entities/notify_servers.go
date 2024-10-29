package entities

import (
	"context"
	"github.com/ascenmmo/tcp-server/pkg/clients/tcpGameServer"
	"github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	"github.com/google/uuid"
)

type NotifyServers interface {
	NotifyServers(ids []uuid.UUID, req types.RequestSetMessage) error
	AddServer(ID uuid.UUID, token, addr string) error
}

type notifier struct {
	servers []*server
}

func NewNotifierServers() NotifyServers {
	return &notifier{}
}

type server struct {
	ID   uuid.UUID `json:"id"`
	Addr string    `json:"addr"`
}

func (n *notifier) NotifyServers(ids []uuid.UUID, req types.RequestSetMessage) error {
	for _, id := range ids {
		for _, server := range n.servers {
			if server.ID == id {
				err := tcpGameServer.New(server.Addr).GameConnections().SetSendMessage(context.Background(), req.Token, req)
				if err != nil {
					n.RemoveNotifyServer(server.ID)
					return err
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
			n.servers = append(n.servers[:i], n.servers[i+1:]...)
		}
	}
}

func (s *server) Connect(token string) error {
	cli := tcpGameServer.New(s.Addr)
	_, err := cli.ServerSettings().HealthCheck(context.Background(), token)
	if err != nil {
		return err
	}
	return nil
}
