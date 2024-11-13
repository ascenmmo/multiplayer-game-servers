package service

import (
	"github.com/ascenmmo/tcp-server/internal/storage"
	utils2 "github.com/ascenmmo/tcp-server/internal/utils"
	"github.com/ascenmmo/tcp-server/pkg/api/types"
	"github.com/ascenmmo/tcp-server/pkg/errors"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type Service interface {
	GetConnectionsNum() (countConn int, exists bool)
	CreateRoom(token string) (err error)

	SetMessage(token string, req types.RequestSetMessage) (err error)
	GetMessages(token string) (msg types.ResponseGetMessage, err error)

	RemoveUser(userID uuid.UUID, reqToken string) (err error)
}

type service struct {
	maxConnections uint64

	storage memoryDB.IMemoryDB

	token tokengenerator.TokenGenerator

	logger zerolog.Logger
}

func (s *service) GetConnectionsNum() (countConn int, exists bool) {
	count := s.storage.CountConnection()

	if uint64(count) >= s.maxConnections {
		return count, false
	}

	return count, true
}

func (s *service) CreateRoom(token string) error {
	clientInfo, err := s.token.ParseToken(token)
	if err != nil {
		return err
	}

	roomKey := utils2.GenerateRoomKey(clientInfo)

	_, ok := s.storage.GetData(roomKey)
	if ok {
		return errors.ErrRoomIsExists
	}

	s.setRoom(clientInfo, &types.Room{
		GameID: clientInfo.GameID,
		RoomID: clientInfo.RoomID,
	})

	return nil
}

func (s *service) SetMessage(token string, msg types.RequestSetMessage) (err error) {
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
		room.SetUser(&types.User{
			ID: clientInfo.UserID,
		})
	}

	//s.setRoom(clientInfo, room)

	return nil
}

func (s *service) GetMessages(token string) (msg types.ResponseGetMessage, err error) {
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
		room.SetUser(&types.User{
			ID: clientInfo.UserID,
		})
		//s.setRoom(clientInfo, room)
	}

	return msg, nil
}

func (s *service) RemoveUser(userID uuid.UUID, reqToken string) (err error) {
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

func (s *service) getRoom(clientInfo tokentype.Info) (room *types.Room, err error) {
	roomKey := utils2.GenerateRoomKey(clientInfo)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		newRoom := &types.Room{
			GameID: clientInfo.GameID,
			RoomID: clientInfo.RoomID,
		}
		roomData = newRoom
		s.setRoom(clientInfo, newRoom)

		return newRoom, nil
	}

	room, ok = roomData.(*types.Room)
	if !ok {
		return room, errors.ErrRoomBadValue
	}

	room.UpdatedAt = time.Now()

	return room, nil
}

func (s *service) setRoom(clientInfo tokentype.Info, room *types.Room) {
	roomKey := utils2.GenerateRoomKey(clientInfo)
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
