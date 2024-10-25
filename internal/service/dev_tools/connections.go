package devtools

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type connections struct {
	gameStorage        storage.GameStorage
	serverStorage      storage.ServersStorage
	roomsStorage       storage.RoomsStorage
	gameConfigsStorage storage.GameConfigsStorage
	token              tokengenerator.TokenGenerator
	logger             *zerolog.Logger
}

func (c *connections) CreateRoom(ctx context.Context, token string, gameID uuid.UUID) (newToken string, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return newToken, err
	}

	game, err := c.gameStorage.FindByID(gameID)
	if err != nil {
		return newToken, err
	}

	room, err := c.roomsStorage.FindByCreatorID(info.UserID)
	if err != nil {
		room = types.Room{
			ID:          uuid.New(),
			GameID:      gameID,
			CreatorID:   info.UserID,
			Connections: []uuid.UUID{info.UserID},
			CreatedAt:   time.Now().Unix(),
		}
	}

	servers, err := c.serverStorage.FindByIDs(game.Servers)
	if err != nil {
		return newToken, errors.ErrServerCreatingRoomOllServesOffError
	}

	for i := range servers {
		exists, err := servers[i].IsExists(ctx, token)
		if err != nil {
			c.logger.Error().Err(err)
			continue
		}
		if exists {
			room.Servers = append(room.Servers, servers[i].ID)
		}
	}

	if len(room.Servers) == 0 {
		return newToken, errors.ErrServerCreatingRoomOllServesOffError
	}

	err = c.roomsStorage.CreateRoom(room)
	if err != nil {
		return "", err
	}

	generateToken, err := c.token.GenerateToken(tokentype.Info{
		GameID: gameID,
		RoomID: room.ID,
		UserID: info.UserID,
	}, tokengenerator.AESGCM)
	if err != nil {
		return newToken, err
	}

	return generateToken, nil
}

func (c *connections) GetRoomsAll(ctx context.Context, token string, gameID uuid.UUID) (rooms []types.Room, err error) {
	_, err = c.token.ParseToken(token)
	if err != nil {
		return rooms, err
	}

	rooms, err = c.roomsStorage.FindAll(gameID)
	if err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (c *connections) JoinRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (newToken string, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return newToken, err
	}

	room, err := c.roomsStorage.FindByID(roomID)
	if err != nil {
		return newToken, err
	}

	room.Connections = append(room.Connections, info.UserID)

	err = c.roomsStorage.Update(room)
	if err != nil {
		return "", err
	}

	generateToken, err := c.token.GenerateToken(tokentype.Info{
		GameID: gameID,
		RoomID: room.ID,
		UserID: info.UserID,
	}, tokengenerator.AESGCM)
	if err != nil {
		return newToken, err
	}

	return generateToken, nil
}

func (c *connections) GetRoomsConnectionUrls(ctx context.Context, token string) (connectionsServer []types.ConnectionServer, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return []types.ConnectionServer{}, err
	}
	if uuid.Nil == info.RoomID || uuid.Nil == info.GameID {
		return []types.ConnectionServer{}, errors.ErrAccessDenied
	}

	room, err := c.roomsStorage.FindByID(info.RoomID)
	if err != nil {
		return []types.ConnectionServer{}, err
	}

	servers, err := c.serverStorage.FindByIDs(room.Servers)
	if err != nil {
		return []types.ConnectionServer{}, err
	}

	config, err := c.gameConfigsStorage.GetConfig(info.GameID)
	if err != nil {
		c.logger.Error().Err(err).Msg("game configs not found")
	}

	for _, server := range servers {
		err = server.CreateRoom(ctx, token, config)
		if err != nil {
			if err.Error() != errors.ErrRoomIsExists.Error() {
				c.logger.Error().Err(err).Msg("server error create room")
				continue
			}
		}
		connectionsServer = append(connectionsServer, server.GetConnectionServer())
	}

	return connectionsServer, nil
}

func (c *connections) RemoveRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return err
	}

	room, err := c.roomsStorage.FindByID(roomID)
	if err != nil {
		return err
	}

	if room.CreatorID != info.UserID {
		return errors.ErrAccessDenied
	}

	err = c.roomsStorage.Delete(room.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewConnections(gameStorage storage.GameStorage, serverStorage storage.ServersStorage, roomsStorage storage.RoomsStorage, gameConfigsStorage storage.GameConfigsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.DevToolsConnections {
	return &connections{gameStorage: gameStorage, serverStorage: serverStorage, roomsStorage: roomsStorage, gameConfigsStorage: gameConfigsStorage, token: token, logger: logger}
}
