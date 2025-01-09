package types

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ascenmmo/udp-server/pkg/api/types"
	"github.com/ascenmmo/udp-server/pkg/clients/udpGameServer"
	"github.com/google/uuid"
	"strings"
	"time"
)

const (
	ServerTypeUDP       = "udp"
	ServerTypeTCP       = "tcp"
	ServerTypeWebsocket = "websocket"
)

type Server struct {
	ID         uuid.UUID `json:"id" bson:"_id"`
	ServerName string    `json:"server_name" bson:"server_name"`

	Owners         []uuid.UUID `json:"-" bson:"owners"`
	Address        string      `json:"address" bson:"address"`
	ServerType     string      `json:"server_type" bson:"server_type"`
	Region         string      `json:"region" bson:"region"`
	ConnectionPort string      `json:"connection_port" bson:"connection_port"`
	MaxConnections int         `json:"maxConnections" bson:"max_connections"`

	AscenmmoServer bool `json:"ascenmmo_server" bson:"ascenmmo_server"`
	IsActive       bool `json:"is_active" bson:"is_active"`
}

type ConnectionServer struct {
	ID             uuid.UUID `json:"id"`
	Address        string    `json:"address"`
	ConnectionPort string    `json:"connectionPort"`
	Path           string    `json:"path"`
	FullURL        string    `json:"fullURL"`
	ServerType     string    `json:"serverType"`
	IsExists       bool      `json:"isExists"`
}

func (s *Server) GetConnectionsNum(ctx context.Context, token string) (countConn int, exists bool, err error) {
	cli := udpGameServer.New(s.getRestUrl())
	countConn, exists, err = cli.ServerSettings().GetConnectionsNum(ctx, token)
	return countConn, exists, err
}

func (s *Server) IsExists(ctx context.Context, token string) (bool, error) {
	var err error
	cli := udpGameServer.New(s.getRestUrl())
	settings, err := cli.ServerSettings().GetServerSettings(ctx, token)
	if err != nil {
		s.IsActive = false
		return false, err
	}

	s.ServerType = settings.ServerType
	s.ConnectionPort = settings.ConnectionPort
	s.IsActive = true
	s.MaxConnections = settings.MaxConnections

	return s.IsActive, err
}

func (s *Server) CreateRoom(ctx context.Context, token string, ttl time.Duration) (err error) {
	cli := udpGameServer.New(s.getRestUrl())

	err = cli.ServerSettings().CreateRoom(ctx, token, types.CreateRoomRequest{
		RoomTTl: ttl,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) GetDeletedRooms(ctx context.Context, token string, ids []GetDeletedRooms) (resultIDs []GetDeletedRooms, err error) {
	cli := udpGameServer.New(s.getRestUrl())

	var idsArray []types.GetDeletedRooms
	marshal, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshal, &idsArray)
	if err != nil {
		return nil, err
	}

	resp, err := cli.ServerSettings().GetDeletedRooms(ctx, token, idsArray)
	if err != nil {
		return nil, err
	}

	m, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, &resultIDs)
	if err != nil {
		return nil, err
	}

	return resultIDs, nil
}

func (s *Server) RemoveOwner(ownerID uuid.UUID) {
	for i, ownersID := range s.Owners {
		if ownersID == ownerID {
			s.Owners = append(s.Owners[:i], s.Owners[i+1:]...)
			break
		}
	}
}

func (s *Server) AddOwner(ownerID uuid.UUID) {
	uniqueOwners := make(map[uuid.UUID]struct{})

	s.Owners = append(s.Owners, ownerID)

	for _, owner := range s.Owners {
		uniqueOwners[owner] = struct{}{}
	}

	newOwners := make([]uuid.UUID, 0, len(uniqueOwners))

	for owner, _ := range uniqueOwners {
		newOwners = append(newOwners, owner)
	}
}

func (s *Server) IsOwner(ownerID uuid.UUID) bool {
	for _, owner := range s.Owners {
		if owner == ownerID {
			return true
		}
	}
	return false
}

func (s *Server) GetConnectionServer() (c ConnectionServer) {
	c.ID = s.ID
	c.Address = s.gatSingleAddress()
	c.ConnectionPort = s.ConnectionPort

	switch s.ServerType {
	case ServerTypeTCP:
		c.Path = "/api/v1/rest/gameConnections/"
	case ServerTypeUDP:
		c.Path = ""
	case ServerTypeWebsocket:
		c.Path = "/api/ws/connect"
	}

	c.ServerType = s.ServerType
	c.FullURL = c.Address + ":" + c.ConnectionPort + c.Path
	c.IsExists = true

	return c
}

func (s *Server) gatSingleAddress() string {
	raplase := strings.Replace(s.Address, "https://", "", 1)
	raplase = strings.Replace(raplase, "http://", "", 1)
	split := strings.Split(raplase, ":")
	return split[0]
}

func (s *Server) getRestUrl() string {
	raplase := strings.Replace(s.Address, "https://", "", 1)
	raplase = strings.Replace(raplase, "http://", "", 1)
	return fmt.Sprintf("http://%s", raplase)
}
