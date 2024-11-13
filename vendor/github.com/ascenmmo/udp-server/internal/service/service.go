package service

import (
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/udp-server/internal/connection"
	memoryDB "github.com/ascenmmo/udp-server/internal/storage"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/ascenmmo/udp-server/pkg/api/types"
	"github.com/ascenmmo/udp-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Service interface {
	GetConnectionsNum() (countConn int, exists bool)
	CreateRoom(token string) error
	GetUsersAndMessage(ds connection.DataSender, req []byte) (users []types.User, msg []byte, err error)
	RemoveUser(ds connection.DataSender, userID uuid.UUID) (err error)
}

type service struct {
	maxConnections uint64

	storage memoryDB.IMemoryDB

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
func (s *service) CreateRoom(token string) error {
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
	})

	return nil
}

func (s *service) GetUsersAndMessage(ds connection.DataSender, req []byte) (users []types.User, msg []byte, err error) {
	clientInfo, room, err := s.getRoom(ds)
	if err != nil {
		clientInfo, err = s.setNewUser(ds, req)
		if err != nil {
			return nil, nil, err
		}
		return append(users, types.User{Connection: ds}), []byte(clientInfo.UserID.String()), nil
	}

	if len(req) == 454 || len(req) == 343 {
		return append(users, types.User{Connection: ds}), []byte(clientInfo.UserID.String()), nil
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

func (s *service) setNewUser(ds connection.DataSender, req []byte) (clientInfo *tokentype.Info, err error) {
	token := string(req)

	info, err := s.token.ParseToken(token)
	if err != nil {
		return clientInfo, errors.ErrNewConnectionMastGetToken
	}
	clientInfo = &info
	s.storage.SetData(ds.GetID(), info)
	roomKey := utils.GenerateRoomKey(info)

	var room *types.Room
	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		room = &types.Room{
			GameID: clientInfo.GameID,
			RoomID: clientInfo.RoomID,
		}
		s.setRoom(*clientInfo, room)
	} else {
		room, ok = roomData.(*types.Room)
		if !ok {
			return clientInfo, errors.ErrRoomBadValue
		}
	}

	room.SetUser(&types.User{
		ID:         info.UserID,
		Connection: ds,
	})

	s.storage.AddConnection(token)

	return clientInfo, nil
}

func (s *service) getRoom(ds connection.DataSender) (clientInfo *tokentype.Info, room *types.Room, err error) {
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

	room, ok = roomData.(*types.Room)
	if !ok {
		return nil, nil, errors.ErrRoomBadValue
	}

	return &info, room, nil
}

func (s *service) getRoomByClientInfo(clientInfo tokentype.Info) (room *types.Room, err error) {
	roomKey := utils.GenerateRoomKey(clientInfo)

	roomData, ok := s.storage.GetData(roomKey)
	if !ok {
		return nil, errors.ErrRoomNotFound
	}

	room, ok = roomData.(*types.Room)
	if !ok {
		return nil, errors.ErrRoomBadValue
	}

	return room, nil
}

func (s *service) setRoom(clientInfo tokentype.Info, room *types.Room) {
	roomKey := utils.GenerateRoomKey(clientInfo)
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
