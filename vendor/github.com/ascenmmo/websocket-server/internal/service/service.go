package service

import (
	"encoding/json"
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/websocket-server/internal/connection"
	configsService "github.com/ascenmmo/websocket-server/internal/service/configs_service"
	"github.com/ascenmmo/websocket-server/internal/storage"
	"github.com/ascenmmo/websocket-server/internal/utils"
	"github.com/ascenmmo/websocket-server/pkg/entities"
	"github.com/ascenmmo/websocket-server/pkg/errors"
	"github.com/ascenmmo/websocket-server/pkg/restconnection/types"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"runtime"
	"time"
)

type Service interface {
	GetConnectionsNum() (countConn int, exists bool)
	CreateRoom(token string, configs types.GameConfigs) error
	SetRoomNotifyServer(token string, id uuid.UUID, url string) (err error)
	NotifyAllServers(clientInfo tokentype.Info, req types.Request) (err error)
	GetUsersAndMessage(ds connection.DataSender, clientInfo tokentype.Info, req types.Request) (users []entities.User, msg []byte, err error)
	RemoveUser(token string, userID uuid.UUID) (err error)
	ParseToken(token string) (info tokentype.Info, err error)
	SetNewConnection(clientInfo tokentype.Info, ds connection.DataSender) (err error)
	GetGameResults(token string) (results []types.GameConfigResults, err error)
}

type service struct {
	maxConnections    uint64
	gameConfigService configsService.GameConfigsService
	storage           memoryDB.IMemoryDB
	token             tokengenerator.TokenGenerator

	logger zerolog.Logger
}

func (s *service) GetConnectionsNum() (countConn int, exists bool) {
	count := s.storage.CountConnection()

	if uint64(count) >= s.maxConnections {
		return count, false
	}

	return count, true
}

func (s *service) SetRoomNotifyServer(token string, id uuid.UUID, url string) (err error) {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	room, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}

	room.SetServerID(id)

	data, _ := s.storage.GetData(utils.GenerateNotifyServerKey())

	server, ok := data.(connection.NotifyServers)
	if !ok {
		s.logger.Warn().Msg("NotifyServers cant get interfase")
		server = connection.NewNotifierServers()
	}

	err = server.AddServer(id, token, url)
	if err != nil {
		return err
	}

	s.storage.SetData(utils.GenerateNotifyServerKey(), server)

	return nil

}

func (s *service) NotifyAllServers(clientInfo tokentype.Info, req types.Request) (err error) {
	room, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}
	if len(room.ServerID) == 0 {
		return nil
	}

	data, ok := s.storage.GetData(utils.GenerateNotifyServerKey())
	if !ok {
		return errors.ErrNotifyServerNotFound
	}

	servers, ok := data.(connection.NotifyServers)
	if !ok {
		return errors.ErrNotifyServerNotValid
	}

	err = servers.NotifyServers(room.ServerID, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateRoom(token string, configs types.GameConfigs) error {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	roomKey := utils.GenerateRoomKey(clientInfo)

	_, ok := s.storage.GetData(roomKey)
	if ok {
		return errors.ErrRoomIsExists
	}

	configs = s.gameConfigService.SetServerExecuteToGameConfig(clientInfo, configs)

	s.setRoom(clientInfo, &entities.Room{
		GameID:      clientInfo.GameID,
		RoomID:      clientInfo.RoomID,
		GameConfigs: configs,
	})

	return nil
}

func (s *service) GetUsersAndMessage(ds connection.DataSender, clientInfo tokentype.Info, req types.Request) (users []entities.User, msg []byte, err error) {
	room, err := s.getRoom(clientInfo)
	if err != nil {
		return nil, nil, err
	}

	isNew := true
	usersData := room.GetUser()
	for _, v := range usersData {
		if v.ID == clientInfo.UserID &&
			ds.GetID() == v.Connection.GetID() {
			isNew = false
			continue
		}

		users = append(users, *v)
	}

	if isNew {
		room.SetUser(&entities.User{
			ID:         clientInfo.UserID,
			Connection: ds,
		})

		s.storage.AddConnection(clientInfo.UserID.String())
	}

	response := types.Response{
		Data: req.Data,
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		return nil, nil, err
	}

	if req.Server == nil {
		s.gameConfigService.Do(req.Token, clientInfo, room.GameConfigs, req.Data)
		id := uuid.New()
		req.Server = &id
		err = s.NotifyAllServers(clientInfo, req)
		if err != nil {
			s.logger.Warn().Err(err).Msg("NotifyAllServers err")
		}
	}

	return users, marshal, err
}

func (s *service) RemoveUser(token string, userID uuid.UUID) (err error) {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	game, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}

	game.RemoveUser(userID)

	return nil
}

func (s *service) ParseToken(token string) (info tokentype.Info, err error) {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return info, err
	}

	return clientInfo, nil
}

func (s *service) SetNewConnection(clientInfo tokentype.Info, ds connection.DataSender) (err error) {
	room, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}
	room.SetUser(&entities.User{
		ID:         clientInfo.UserID,
		Connection: ds,
	})

	s.setRoom(clientInfo, room)

	return nil
}

func (s *service) GetGameResults(token string) (results []types.GameConfigResults, err error) {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return results, err
	}

	playersOnline := s.storage.GetAllConnection()
	roomsResults, ok := s.gameConfigService.GetDeletedRoomsResults(clientInfo, playersOnline)
	if !ok {
		return results, errors.ErrGameResultsNotFound
	}

	return roomsResults, nil
}

func (s *service) getRoom(clientInfo tokentype.Info) (room *entities.Room, err error) {
	roomKey := utils.GenerateRoomKey(clientInfo)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		return room, errors.ErrRoomNotFound
	}

	room, ok = roomData.(*entities.Room)
	if !ok {
		return room, errors.ErrRoomBadValue
	}

	return room, nil
}

func (s *service) setRoom(clientInfo tokentype.Info, room *entities.Room) {
	roomKey := utils.GenerateRoomKey(clientInfo)
	s.storage.SetData(roomKey, room)
}

func NewService(token tokengenerator.TokenGenerator, storage memoryDB.IMemoryDB, gameConfigService configsService.GameConfigsService, logger zerolog.Logger) Service {
	srv := &service{
		maxConnections:    uint64(types.CountConnectionsMAX()),
		storage:           storage,
		token:             token,
		gameConfigService: gameConfigService,
		logger:            logger,
	}
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		for range ticker.C {
			fmt.Println(fmt.Sprintf("count connections: %d \t max conections: %d", srv.storage.CountConnection(), srv.maxConnections))
			fmt.Println(fmt.Sprintf("count gorutines: %d ", runtime.NumGoroutine()))
		}
	}()
	return srv
}
