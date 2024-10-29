package types

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ascenmmo/tcp-server/pkg/clients/tcpGameServer"
	restType "github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	"github.com/ascenmmo/udp-server/pkg/clients/udpGameServer"
	"github.com/google/uuid"
	"strings"
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

	AscenmmoServer bool `json:"ascenmmo_server" bson:"ascenmmo_server"`
	IsActive       bool `json:"is_active" bson:"is_active"`
}

type ConnectionServer struct {
	Address        string
	ConnectionPort string
	Path           string
	FullURL        string
	ServerType     string
}

func (s *Server) IsExists(ctx context.Context, token string) (bool, error) {
	var err error
	cli := udpGameServer.New(s.getRestUrl())
	exists, err := cli.ServerSettings().HealthCheck(context.TODO(), token)
	if err != nil {
		s.IsActive = false
		return false, err
	}

	settings, err := cli.ServerSettings().GetServerSettings(ctx, token)
	if err != nil {
		s.IsActive = false
		return false, err
	}

	s.ServerType = settings.ServerType
	s.ConnectionPort = settings.ConnectionPort
	s.IsActive = true

	s.IsActive = exists
	return exists, err
}

func (s *Server) CreateRoom(ctx context.Context, token string, gameConfigs GameConfigs) (err error) {
	cli := tcpGameServer.New(s.getRestUrl())

	confBuf, err := json.Marshal(gameConfigs)

	var config restType.GameConfigs
	err = json.Unmarshal(confBuf, &config)
	if err != nil {
		return err
	}

	err = cli.ServerSettings().CreateRoom(ctx, token, restType.CreateRoomRequest{
		GameConfigs: config,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) GetGameConfigResults(ctx context.Context, token string) (gameResults []GameConfigResults, err error) {
	cli := tcpGameServer.New(s.getRestUrl())
	results, err := cli.ServerSettings().GetGameResults(ctx, token)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		gameResults = append(gameResults, GameConfigResults{
			ID:     uuid.New(),
			GameID: v.GameID,
			RoomID: v.RoomID,
			Result: v.Result,
		})
	}
	return gameResults, err
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
