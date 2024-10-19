package types

import (
	"context"
	"github.com/ascenmmo/tcp-server/pkg/clients/tcpGameServer"
	restType "github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	"github.com/ascenmmo/udp-server/pkg/clients/udpGameServer"
	udpType "github.com/ascenmmo/udp-server/pkg/restconnection/types"
	"github.com/ascenmmo/websocket-server/pkg/clients/wsGameServer"
	wsType "github.com/ascenmmo/websocket-server/pkg/restconnection/types"
	"github.com/google/uuid"
)

const (
	ServerTypeUDP       = "udp"
	ServerTypeTCP       = "tcp"
	ServerTypeWebsocket = "websocket"
	ServerTypeChat      = "chat"
)

type Server struct {
	ID         uuid.UUID `json:"id" bson:"_id"`
	ServerName string    `json:"server_name" bson:"server_name"`

	Owners     []uuid.UUID `json:"-" bson:"owners"`
	Address    string      `json:"address" bson:"address"`
	ServerType string      `json:"server_type" bson:"server_type"`
	Region     string      `json:"region" bson:"region"`

	AscenmmoServer bool `json:"ascenmmo_server" bson:"ascenmmo_server"`
	IsActive       bool `json:"is_active" bson:"is_active"`
}

type ConnectionServer struct {
	Address    string
	ServerType string
}

func (s *Server) IsExists(ctx context.Context, token string) bool {
	var err error
	switch s.ServerType {
	case ServerTypeUDP:
		err = s.isExistsUDP(ctx, token)
	case ServerTypeTCP:
		err = s.isExistsRest(ctx, token)
	case ServerTypeWebsocket:
		err = s.isExistsWS(ctx, token)
	case ServerTypeChat:
		err = s.isExistsWS(ctx, token)
	default:
		return false
	}
	return err == nil
}

func (s *Server) CreateRoom(ctx context.Context, token string, gameConfigs GameConfigs) (err error) {
	switch s.ServerType {
	case ServerTypeUDP:
		err = s.createUDPRoom(ctx, token)
	case ServerTypeTCP:
		err = s.createRestRoom(ctx, token)
	case ServerTypeWebsocket:
		err = s.createWSRoom(ctx, token)
	}
	return err
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

func (s *Server) createRestRoom(ctx context.Context, token string) (err error) {
	cli := tcpGameServer.New(s.Address)

	err = cli.ServerSettings().CreateRoom(ctx, token, restType.CreateRoomRequest{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) createUDPRoom(ctx context.Context, token string) (err error) {
	cli := udpGameServer.New(s.Address)
	err = cli.ServerSettings().CreateRoom(ctx, token, udpType.CreateRoomRequest{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) createWSRoom(ctx context.Context, token string) (err error) {
	cli := wsGameServer.New(s.Address)
	err = cli.ServerSettings().CreateRoom(ctx, token, wsType.CreateRoomRequest{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) isExistsWS(ctx context.Context, token string) (err error) {
	cli := wsGameServer.New(s.Address)
	exists, err := cli.ServerSettings().HealthCheck(context.TODO(), token)
	if err != nil {
		s.IsActive = false
		return err
	}

	settings, err := cli.ServerSettings().GetServerSettings(ctx, token)
	if err != nil {
		s.IsActive = false
		return err
	}

	s.ServerType = settings.ServerType
	s.IsActive = true

	s.IsActive = exists
	return nil
}

func (s *Server) isExistsUDP(ctx context.Context, token string) (err error) {
	cli := udpGameServer.New(s.Address)
	exists, err := cli.ServerSettings().HealthCheck(context.TODO(), token)
	if err != nil {
		s.IsActive = false
		return err
	}

	settings, err := cli.ServerSettings().GetServerSettings(ctx, token)
	if err != nil {
		s.IsActive = false
		return err
	}

	s.ServerType = settings.ServerType
	s.IsActive = true

	s.IsActive = exists
	return nil
}

func (s *Server) isExistsRest(ctx context.Context, token string) (err error) {
	cli := tcpGameServer.New(s.Address)
	exists, err := cli.ServerSettings().HealthCheck(context.TODO(), token)
	if err != nil {
		s.IsActive = false
		return err
	}

	settings, err := cli.ServerSettings().GetServerSettings(ctx, token)
	if err != nil {
		s.IsActive = false
		return err
	}

	s.ServerType = settings.ServerType
	s.IsActive = true

	s.IsActive = exists
	return nil
}
