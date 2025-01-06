package service

import (
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/websocket-server/internal/connection"
	"github.com/ascenmmo/websocket-server/internal/storage"
	"github.com/ascenmmo/websocket-server/internal/utils"
	"github.com/ascenmmo/websocket-server/pkg/api/types"
	"github.com/ascenmmo/websocket-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type Service interface {
	GetConnectionsNum() (countConn int, exists bool)
	CreateRoom(token string, request types.CreateRoomRequest) error
	GetUsersAndMessage(ds connection.DataSender, clientInfo tokentype.Info, req []byte) (users []types.User, msg []byte, err error)
	RemoveUser(clientInfo tokentype.Info, userID uuid.UUID) (err error)
	ParseToken(token string) (info tokentype.Info, err error)
	SetNewConnection(clientInfo tokentype.Info, ds connection.DataSender) (err error)
	GetDeletedRooms(token string, ids []types.GetDeletedRooms) (deletedIds []types.GetDeletedRooms, err error)
}

type service struct {
	maxConnections uint64
	storage        memoryDB.IMemoryDB
	token          tokengenerator.TokenGenerator

	logger zerolog.Logger
}

func (s *service) GetConnectionsNum() (countConn int, exists bool) {
	count := s.storage.CountConnection()

	if uint64(count) >= s.maxConnections {
		return count, false
	}

	return count, true
}

func (s *service) CreateRoom(token string, request types.CreateRoomRequest) error {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	roomKey := utils.GenerateRoomKey(clientInfo)

	_, ok := s.storage.GetData(roomKey)
	if ok {
		return errors.ErrRoomIsExists
	}

	s.setRoom(clientInfo, &types.Room{
		GameID: clientInfo.GameID,
		RoomID: clientInfo.RoomID,
	}, request.RoomTTl)

	return nil
}

func (s *service) GetUsersAndMessage(ds connection.DataSender, clientInfo tokentype.Info, req []byte) (users []types.User, msg []byte, err error) {
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
		room.SetUser(&types.User{
			ID:         clientInfo.UserID,
			Connection: ds,
		})

		s.storage.AddConnection(clientInfo.UserID.String())
	}

	return users, req, err
}

func (s *service) RemoveUser(clientInfo tokentype.Info, userID uuid.UUID) (err error) {
	game, err := s.getRoom(clientInfo)
	if err != nil {
		return err
	}

	for _, user := range game.GetUser() {
		if user.ID == userID {
			if user.Connection != nil {
				user.Connection.Close()
			}
		}
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
	room.SetUser(&types.User{
		ID:         clientInfo.UserID,
		Connection: ds,
	})

	return nil
}

func (s *service) GetDeletedRooms(token string, ids []types.GetDeletedRooms) (deletedIds []types.GetDeletedRooms, err error) {
	info, err := s.token.ParseToken(token)
	if err != nil {
		return nil, err
	}

	roomsWithKey := make(map[string]types.GetDeletedRooms)
	for _, id := range ids {
		info.GameID = id.GameID
		info.RoomID = id.RoomID
		roomsWithKey[utils.GenerateRoomKey(info)] = id
	}

	for k, _ := range roomsWithKey {
		_, ok := s.storage.GetData(k)
		if !ok {
			delete(roomsWithKey, k)
		}
	}

	for _, v := range roomsWithKey {
		deletedIds = append(deletedIds, v)
	}

	return deletedIds, nil
}

func (s *service) getRoom(clientInfo tokentype.Info) (room *types.Room, err error) {
	roomKey := utils.GenerateRoomKey(clientInfo)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		newRoom := &types.Room{
			GameID: clientInfo.GameID,
			RoomID: clientInfo.RoomID,
		}
		s.setRoom(clientInfo, newRoom, 0)
		roomData = newRoom
	}

	room, ok = roomData.(*types.Room)
	if !ok {
		return room, errors.ErrRoomBadValue
	}

	return room, nil
}

func (s *service) setRoom(clientInfo tokentype.Info, room *types.Room, ttl time.Duration) {
	roomKey := utils.GenerateRoomKey(clientInfo)
	if ttl != 0 {
		s.storage.SetDataWithTTL(roomKey, room, ttl)
		return
	}
	s.storage.SetData(roomKey, room)
}

func NewService(token tokengenerator.TokenGenerator, storage memoryDB.IMemoryDB, logger zerolog.Logger) Service {
	srv := &service{
		maxConnections: uint64(types.CountConnectionsMAX()),
		storage:        storage,
		token:          token,
		logger:         logger,
	}
	return srv
}
