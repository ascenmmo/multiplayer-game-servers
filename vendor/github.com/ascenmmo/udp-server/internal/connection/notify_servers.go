package connection

import (
	"encoding/json"
	"github.com/google/uuid"
	"net"
)

type NotifyServers interface {
	NotifyServers(ids []uuid.UUID, request []byte) error
	AddServer(ID uuid.UUID, addr string) error
}

type notifier struct {
	servers []*server
}

func NewNotifierServers() NotifyServers {
	return &notifier{}
}

type server struct {
	ID         uuid.UUID `json:"id"`
	Addr       string    `json:"addr"`
	Connection *net.UDPConn
	Add        *net.UDPConn
}

func (n *notifier) NotifyServers(ids []uuid.UUID, request []byte) error {
	if len(n.servers) == 0 {
		return nil
	}
	marshal, err := json.Marshal(request)
	if err != nil {
		return err
	}
	for _, id := range ids {
		for i, server := range n.servers {
			if server.ID == id {
				_, err = n.servers[i].Connection.Write(marshal)
				if err != nil {
					err = n.servers[i].Connect()
					if err != nil {
						n.RemoveNotifyServer(id)
						return err
					}
					_, err = n.servers[i].Connection.Write(marshal)
					return err
				}
			}
		}
	}
	return nil
}

func (n *notifier) AddServer(ID uuid.UUID, addr string) error {
	newServer := &server{
		ID:   ID,
		Addr: addr,
	}
	err := newServer.Connect()
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

func (s *server) Connect() error {
	serverAddr, err := net.ResolveUDPAddr("udp", s.Addr)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return err
	}

	s.Connection = conn

	return nil
}
