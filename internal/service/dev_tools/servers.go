package devtools

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	"github.com/ascenmmo/multiplayer-game-servers/internal/service/access"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	defalultdata "github.com/ascenmmo/multiplayer-game-servers/internal/storage/defalult_data"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type serverService struct {
	access        access.AccessGame
	gameStorage   storage.GameStorage
	serverStorage storage.ServersStorage
	token         tokengenerator.TokenGenerator

	logger *zerolog.Logger
}

func (s *serverService) AddServer(ctx context.Context, token string, name string, address string) (err error) {
	info, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	newServer := &types.Server{
		ID:         uuid.New(),
		ServerName: name,
		Address:    address,
	}

	newServer.AddOwner(info.UserID)

	exists := newServer.IsExists(ctx, token)
	if !exists {
		return errors.ErrServerNotExists
	}

	err = s.serverStorage.CreateServer(*newServer)
	if err != nil {
		return err
	}

	return nil
}

func (s *serverService) GetServers(ctx context.Context, token string) (servers []types.Server, err error) {
	info, err := s.token.ParseToken(token)
	if err != nil {
		return nil, err
	}

	servers = defalultdata.AddServers(storage.DefaultUserID)

	gameServers, err := s.serverStorage.FindAll(info.UserID)
	if err != nil {
		s.logger.Error().Err(err).Msg("find all servers failed")
		return servers, nil
	}

	gameServers = append(gameServers, servers...)

	return gameServers, nil

}

func (s *serverService) DeleteServers(ctx context.Context, token string, serverID uuid.UUID) (err error) {
	info, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	server, err := s.serverStorage.FindByID(serverID)
	if err != nil {
		return err
	}

	if !server.IsOwner(info.UserID) {
		return errors.ErrAccessDenied
	}

	games, err := s.gameStorage.FindByOwnerID(info.UserID)
	if err != nil {
		return err
	}

	for _, game := range games {
		game.RemoveServer(serverID)
		err = s.gameStorage.Update(game)
		if err != nil {
			return err
		}
	}

	err = s.serverStorage.Delete(serverID)
	if err != nil {
		return err
	}

	return nil
}

func NewServerService(access access.AccessGame, gameStorage storage.GameStorage, serverStorage storage.ServersStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.DevToolsServer {
	return &serverService{token: token, access: access, gameStorage: gameStorage, serverStorage: serverStorage, logger: logger}
}
