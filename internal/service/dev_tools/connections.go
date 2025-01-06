package devtools

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/internal/utills"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type SchedulerConnections interface {
	RoomCleaner(ctx context.Context, ttl time.Duration, limit int64)
}

type connections struct {
	gameStorage   storage.GameStorage
	serverStorage storage.ServersStorage
	roomsStorage  storage.RoomsStorage
	token         tokengenerator.TokenGenerator
	logger        *zerolog.Logger
}

func (c *connections) CreateRoom(ctx context.Context, token string, name string) (newToken string, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return newToken, err
	}

	game, err := c.gameStorage.FindByID(info.GameID)
	if err != nil {
		return newToken, errors.ErrGameNotFound
	}

	room, err := c.roomsStorage.FindByCreatorID(info.UserID)
	if err == nil {
		err = c.roomsStorage.Delete(room.ID)
		if err != nil {
			return "", err
		}
	}

	room = types.Room{
		ID:          uuid.New(),
		Name:        name,
		GameID:      info.GameID,
		CreatorID:   info.UserID,
		Connections: []uuid.UUID{info.UserID},
		CreatedAt:   time.Now().Unix(),
	}
	room.RoomCode = utills.GenerateRoomCode(room.ID)

	servers, err := c.serverStorage.FindByIDs(game.Servers)
	if err != nil {
		return newToken, errors.ErrServerCreatingRoomAllServesOffError
	}

	for i := range servers {
		exists, err := servers[i].IsExists(ctx, token)
		if err != nil {
			c.logger.Error().Err(err).Msg("IsExists error")
			continue
		}
		if exists {
			room.Servers = append(room.Servers, servers[i].ID)
		}
	}

	if len(room.Servers) == 0 {
		return newToken, errors.ErrServerCreatingRoomAllServesOffError
	}

	err = c.roomsStorage.CreateRoom(room)
	if err != nil {
		return "", err
	}

	generateToken, err := c.token.GenerateToken(tokentype.Info{
		GameID: info.GameID,
		RoomID: room.ID,
		UserID: info.UserID,
	}, tokengenerator.AESGCM)
	if err != nil {
		return newToken, err
	}

	return generateToken, nil
}

func (c *connections) GetRoomsAll(ctx context.Context, token string) (rooms []types.Room, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return rooms, err
	}

	rooms, err = c.roomsStorage.FindAll(info.GameID)
	if err != nil {
		return rooms, errors.ErrRoomNoFound
	}

	return rooms, nil
}

func (c *connections) JoinRoomByID(ctx context.Context, token string, roomID uuid.UUID) (newToken string, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return newToken, err
	}

	room, err := c.roomsStorage.FindByID(roomID)
	if err != nil {
		return newToken, errors.ErrRoomNoFound
	}

	room.Connections = append(room.Connections, info.UserID)

	err = c.roomsStorage.Update(room)
	if err != nil {
		return "", errors.ErrRoomNoFound
	}

	generateToken, err := c.token.GenerateToken(tokentype.Info{
		GameID: info.GameID,
		RoomID: room.ID,
		UserID: info.UserID,
	}, tokengenerator.AESGCM)
	if err != nil {
		return newToken, err
	}

	return generateToken, nil
}

func (c *connections) JoinRoomByRoomCode(ctx context.Context, token string, roomCode string) (newToken string, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return newToken, err
	}

	room, err := c.roomsStorage.FindByRoomCode(info.GameID, roomCode)
	if err != nil {
		return newToken, errors.ErrRoomNoFound
	}

	room.Connections = append(room.Connections, info.UserID)

	err = c.roomsStorage.Update(room)
	if err != nil {
		return "", errors.ErrRoomNoFound
	}

	generateToken, err := c.token.GenerateToken(tokentype.Info{
		GameID: info.GameID,
		RoomID: room.ID,
		UserID: info.UserID,
	}, tokengenerator.AESGCM)
	if err != nil {
		return newToken, err
	}

	return generateToken, nil
}

func (c *connections) GetMyRoom(ctx context.Context, token string) (room types.Room, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return room, err
	}

	room, err = c.roomsStorage.FindByID(info.RoomID)
	if err != nil {
		return room, errors.ErrRoomNoFound
	}

	return room, nil
}

func (c *connections) LeaveRoom(ctx context.Context, token string, roomID uuid.UUID) (err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return err
	}

	room, err := c.roomsStorage.FindByID(roomID)
	if err != nil {
		return errors.ErrRoomNoFound
	}

	room.RemoveConnectionID(info.UserID)

	err = c.roomsStorage.Update(room)
	if err != nil {
		return errors.ErrRoomNoFound
	}

	return nil
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
		return []types.ConnectionServer{}, errors.ErrRoomNoFound
	}

	servers, err := c.serverStorage.FindByIDs(room.Servers)
	if err != nil {
		return []types.ConnectionServer{}, errors.ErrServerCreatingRoomAllServesOffError
	}

	defer func() {
		if err != nil {
			return
		}
		if len(room.ExistsServers) != 0 {
			for _, server := range servers {
				for _, existServer := range room.ExistsServers {
					if server.ID == existServer {
						connectionsServer = append(connectionsServer, server.GetConnectionServer())
					}
				}
			}
		}
	}()

	uniqueServerType := make(map[string]types.ConnectionServer)
	uniqueAllTypes := make(map[string]struct{})
	for _, server := range servers {
		connNum, exists, err := server.GetConnectionsNum(ctx, token)
		if err != nil {
			c.logger.Error().Err(err).Msg("server error create room")
			continue
		}

		if server.MaxConnections < connNum+len(room.Connections) && exists {
			continue
		}

		if _, ok := uniqueServerType[server.ServerType]; ok {
			continue
		}

		err = server.CreateRoom(ctx, token, time.Duration(time.Minute*10))
		if err != nil {
			if err.Error() != errors.ErrRoomIsExists.Error() {
				c.logger.Error().Err(err).Msg("server error create room")
				continue
			}
		}

		uniqueServerType[server.ServerType] = server.GetConnectionServer()
		uniqueAllTypes[server.ServerType] = struct{}{}
	}

	for key, _ := range uniqueAllTypes {
		if _, ok := uniqueServerType[key]; !ok {
			return connectionsServer, errors.ErrServerFullOfConnections
		}
	}

	for _, server := range uniqueServerType {
		room.ExistsServers = append(room.ExistsServers, server.ID)
	}

	err = c.roomsStorage.Update(room)
	if err != nil {
		return connectionsServer, err
	}

	return connectionsServer, nil
}

func (c *connections) RemoveRoomByID(ctx context.Context, token string, roomID uuid.UUID) (err error) {
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

func (c *connections) RoomCleaner(ctx context.Context, ttl time.Duration, limit int64) {
	ticker := time.NewTicker(ttl)

	for range ticker.C {
		c.logger.Info().Msg("Start Room Cleaner")

		rooms, err := c.roomsStorage.GetByTime(time.Now().Add(time.Minute*15), limit)
		if err != nil {
			continue
		}

		mapServersRoom := make(map[uuid.UUID][]types.GetDeletedRooms)
		mapServer := make(map[uuid.UUID]types.Server)
		for _, room := range rooms {
			servers, err := c.serverStorage.FindByIDs(room.ExistsServers)
			if err != nil {
				continue
			}

			for _, server := range servers {
				mapServer[server.ID] = server
				mapServersRoom[server.ID] = append(mapServersRoom[server.ID], types.GetDeletedRooms{
					GameID: room.GameID,
					RoomID: room.ID,
				})
			}
		}

		token, err := c.token.GenerateToken(tokentype.Info{
			GameID: uuid.New(),
			RoomID: uuid.New(),
		}, tokengenerator.AESGCM)
		if err != nil {
			continue
		}

		deleteRoomsIDs := make(map[uuid.UUID]bool)
		for serverID, roomIDs := range mapServersRoom {
			if server, ok := mapServer[serverID]; ok {
				for _, roomID := range roomIDs {
					deleteRoomsIDs[roomID.RoomID] = false
				}

				deletedRooms, err := server.GetDeletedRooms(ctx, token, roomIDs)
				if err != nil {
					continue
				}

				for _, roomDeleted := range deletedRooms {
					deleteRoomsIDs[roomDeleted.RoomID] = true
				}
			}
		}

		for k, v := range deleteRoomsIDs {
			if v {
				err := c.roomsStorage.Delete(k)
				if err != nil {
					continue
				}
			}
		}

	}
}

func NewConnections(gameStorage storage.GameStorage, serverStorage storage.ServersStorage, roomsStorage storage.RoomsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.DevToolsConnections {
	return &connections{gameStorage: gameStorage, serverStorage: serverStorage, roomsStorage: roomsStorage, token: token, logger: logger}
}

func NewConnectionsScheduler(gameStorage storage.GameStorage, serverStorage storage.ServersStorage, roomsStorage storage.RoomsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) SchedulerConnections {
	return &connections{gameStorage: gameStorage, serverStorage: serverStorage, roomsStorage: roomsStorage, token: token, logger: logger}
}
