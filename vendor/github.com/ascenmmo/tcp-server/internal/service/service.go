package service

import (
	"fmt"
	entities2 "github.com/ascenmmo/tcp-server/internal/entities"
	configsService "github.com/ascenmmo/tcp-server/internal/service/configs_service"
	"github.com/ascenmmo/tcp-server/internal/storage"
	utils2 "github.com/ascenmmo/tcp-server/internal/utils"
	"github.com/ascenmmo/tcp-server/pkg/errors"
	"github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"runtime"
	"sync"
	"time"
)

type Service interface {
	GetConnectionsNum() (countConn int, exists bool)
	CreateRoom(token string, configs types.GameConfigs) (err error)

	SetMessage(token string, req types.RequestSetMessage) (err error)
	GetMessages(token string) (msg types.ResponseGetMessage, err error)

	RemoveUser(userID uuid.UUID, reqToken string) (err error)

	SetRoomNotifyServer(token string, id uuid.UUID, url string) (err error)
	NotifyAllServers(clientInfo tokentype.Info, req types.RequestSetMessage) (err error)
	GetGameResults(token string) (results []types.GameConfigResults, err error)
}

type service struct {
	maxConnections uint64

	storage           memoryDB.IMemoryDB
	gameConfigService configsService.GameConfigsService

	token tokengenerator.TokenGenerator

	logger zerolog.Logger
	mtx    sync.Mutex
}

func (s *service) GetConnectionsNum() (countConn int, exists bool) {
	count := s.storage.CountConnection()

	if uint64(count) >= s.maxConnections {
		return count, false
	}

	return count, true
}

func (s *service) CreateRoom(token string, configs types.GameConfigs) error {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	roomKey := utils2.GenerateRoomKey(clientInfo)

	_, ok := s.storage.GetData(roomKey)
	if ok {
		return errors.ErrRoomIsExists
	}

	configs = s.gameConfigService.SetServerExecuteToGameConfig(clientInfo, configs)

	s.setRoom(clientInfo, &entities2.Room{
		GameID:      clientInfo.GameID,
		RoomID:      clientInfo.RoomID,
		GameConfigs: configs,
	})

	return nil
}

func (s *service) SetMessage(token string, msg types.RequestSetMessage) (err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	room, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}

	isFound := false
	for i, user := range room.Users {
		if user.ID == clientInfo.UserID {
			isFound = true
			continue
		}

		userData := room.Users[i].Data
		userData = append(userData, msg.Data)
		room.Users[i].Data = userData
	}

	if !isFound {
		room.SetUser(&entities2.User{
			ID: clientInfo.UserID,
		})
	}

	s.setRoom(clientInfo, room)

	if msg.Server == nil {
		s.gameConfigService.Do(token, clientInfo, room.GameConfigs, msg.Data)
		id := uuid.New()
		msg.Server = &id
		msg.Token = token
		err := s.NotifyAllServers(clientInfo, msg)
		if err != nil {
			s.logger.Warn().Err(err).Msg("failed to notify servers")
		}
	}

	return nil
}

func (s *service) GetMessages(token string) (msg types.ResponseGetMessage, err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return msg, err
	}

	room, err := s.getRoom(clientInfo)
	if err != nil {
		return msg, err
	}

	isFound := false
	for i, user := range room.Users {
		if user.ID == clientInfo.UserID {
			isFound = true
			msg.DataArray = user.Data
			room.Users[i].Data = nil
			return msg, nil
		}
	}

	if !isFound {
		room.SetUser(&entities2.User{
			ID: clientInfo.UserID,
		})
		//s.setRoom(clientInfo, room)
	}

	return msg, nil
}

func (s *service) RemoveUser(userID uuid.UUID, reqToken string) (err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	clientInfo, err := s.token.ParseToken(reqToken)
	if err != nil {
		return err
	}

	room, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}

	room.RemoveUser(userID)

	s.setRoom(clientInfo, room)
	return nil
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

	data, _ := s.storage.GetData(utils2.GenerateNotifyServerKey())

	server, ok := data.(entities2.NotifyServers)
	if !ok {
		s.logger.Warn().Msg("NotifyServers cant get interfase")
		server = entities2.NewNotifierServers()
	}

	err = server.AddServer(id, token, url)
	if err != nil {
		return err
	}

	s.storage.SetData(utils2.GenerateNotifyServerKey(), server)

	return nil

}

func (s *service) NotifyAllServers(clientInfo tokentype.Info, req types.RequestSetMessage) (err error) {
	room, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}
	if len(room.ServerID) == 0 {
		return nil
	}

	data, ok := s.storage.GetData(utils2.GenerateNotifyServerKey())
	if !ok {
		return errors.ErrNotifyServerNotFound
	}

	servers, ok := data.(entities2.NotifyServers)
	if !ok {
		return errors.ErrNotifyServerNotValid
	}

	err = servers.NotifyServers(room.ServerID, req)
	if err != nil {
		return err
	}

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

func (s *service) getRoom(clientInfo tokentype.Info) (room *entities2.Room, err error) {
	roomKey := utils2.GenerateRoomKey(clientInfo)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		return room, errors.ErrRoomNotFound
	}

	room, ok = roomData.(*entities2.Room)
	if !ok {
		return room, errors.ErrRoomBadValue
	}

	room.UpdatedAt = time.Now()

	return room, nil
}

func (s *service) setRoom(clientInfo tokentype.Info, room *entities2.Room) {
	roomKey := utils2.GenerateRoomKey(clientInfo)
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
