package service

import (
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/udp-server/internal/connection"
	"github.com/ascenmmo/udp-server/internal/entities"
	configsService "github.com/ascenmmo/udp-server/internal/service/configs_service"
	memoryDB "github.com/ascenmmo/udp-server/internal/storage"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/ascenmmo/udp-server/pkg/errors"
	"github.com/ascenmmo/udp-server/pkg/restconnection/types"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type Service interface {
	GetConnectionsNum() (countConn int, exists bool)
	CreateRoom(token string, configs types.GameConfigs) error
	GetUsersAndMessage(ds connection.DataSender, req []byte) (users []entities.User, msg []byte, err error)
	NotifyAllServers(clientInfo tokentype.Info, reqreq []byte) (err error)
	RemoveUser(ds connection.DataSender, userID uuid.UUID) (err error)
	SetRoomNotifyServer(token string, id uuid.UUID, url string) (err error)
	GetGameResults(token string) (results []types.GameConfigResults, err error)
}

type service struct {
	maxConnections uint64

	storage           memoryDB.IMemoryDB
	gameConfigService configsService.GameConfigsService

	token  tokengenerator.TokenGenerator
	logger zerolog.Logger
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

func (s *service) SetRoomNotifyServer(token string, id uuid.UUID, url string) (err error) {
	//clientInfo, err := s.token.ParseToken(token)
	//if err != nil {
	//	return err
	//}
	//
	//room, err := s.getRoom(clientInfo)
	//if err != nil {
	//	return err
	//}
	//
	//room.SetServerID(id)
	//
	//data, _ := s.storage.GetData(utils.GenerateNotifyServerKey())
	//
	//server, ok := data.(connection.NotifyServers)
	//if !ok {
	//	s.logger.Warn().Msg("NotifyServers cant get interfase")
	//	server = connection.NewNotifierServers()
	//}
	//
	//err = server.AddServer(id, url)
	//if err != nil {
	//	return err
	//}
	//
	//s.storage.SetData(utils.GenerateNotifyServerKey(), server)

	return nil
}

func (s *service) NotifyAllServers(clientInfo tokentype.Info, request []byte) (err error) {
	//room, err := s.getRoom(clientInfo)
	//if err != nil {
	//	return err
	//}
	//if len(room.ServerID) == 0 {
	//	return nil
	//}
	//
	//data, ok := s.storage.GetData(utils.GenerateNotifyServerKey())
	//if !ok {
	//	return errors.ErrNotifyServerNotFound
	//}
	//
	//servers, ok := data.(connection.NotifyServers)
	//if !ok {
	//	return errors.ErrNotifyServerNotValid
	//}
	//
	//err = servers.NotifyServers(room.ServerID, request)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s *service) GetUsersAndMessage(ds connection.DataSender, req []byte) (users []entities.User, msg []byte, err error) {
	clientInfo, room, err := s.getRoom(ds)
	if err != nil {
		clientInfo, err = s.setNewUser(ds, req)
		if err != nil {
			return nil, nil, err
		}
		return append(users, entities.User{Connection: ds}), []byte(clientInfo.UserID.String()), nil
	}

	if len(req) == 454 || len(req) == 343 {
		return append(users, entities.User{Connection: ds}), []byte(clientInfo.UserID.String()), nil
	}

	usersData := room.GetUser()
	for _, v := range usersData {
		if v.ID == clientInfo.UserID &&
			ds.GetID() == v.Connection.GetID() {
			continue
		}
		users = append(users, *v)
	}

	msg = req

	return users, msg, err
}

func (s *service) RemoveUser(ds connection.DataSender, userID uuid.UUID) (err error) {
	_, room, err := s.getRoom(ds)
	if err != nil {
		return err
	}

	room.RemoveUser(userID)

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

func (s *service) setNewUser(ds connection.DataSender, req []byte) (clientInfo *tokentype.Info, err error) {
	token := string(req)

	info, err := s.token.ParseToken(token)
	if err != nil {
		return clientInfo, errors.ErrNewConnectionMastGetToken
	}
	clientInfo = &info
	s.storage.SetData(ds.GetID(), info)
	roomKey := utils.GenerateRoomKey(info)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		return clientInfo, errors.ErrRoomNotFound
	}

	room, ok := roomData.(*entities.Room)
	if !ok {
		return clientInfo, errors.ErrRoomBadValue
	}

	room.SetUser(&entities.User{
		ID:         info.UserID,
		Connection: ds,
	})

	s.storage.AddConnection(token)

	return clientInfo, nil
}

func (s *service) getRoom(ds connection.DataSender) (clientInfo *tokentype.Info, room *entities.Room, err error) {
	client, ok := s.storage.GetData(ds.GetID())
	if !ok {
		return nil, nil, errors.ErrUserNotFound
	}

	info, ok := client.(tokentype.Info)
	if !ok {
		return nil, nil, errors.ErrUserBadValue
	}

	roomKey := utils.GenerateRoomKey(info)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		return nil, nil, errors.ErrRoomNotFound
	}

	room, ok = roomData.(*entities.Room)
	if !ok {
		return nil, nil, errors.ErrRoomBadValue
	}

	return &info, room, nil
}

func (s *service) getRoomByClientInfo(clientInfo tokentype.Info) (room *entities.Room, err error) {
	roomKey := utils.GenerateRoomKey(clientInfo)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		return nil, errors.ErrRoomNotFound
	}

	room, ok = roomData.(*entities.Room)
	if !ok {
		return nil, errors.ErrRoomBadValue
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
			logger.Info().Msg(fmt.Sprintf("count connections: %d \t max conections: %d", srv.storage.CountConnection(), srv.maxConnections))
		}
	}()
	return srv
}
