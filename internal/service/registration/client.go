package registration

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type clientService struct {
	clientStorage storage.ClientStorage
	gameStorage   storage.GameStorage
	roomStorage   storage.RoomsStorage
	token         tokengenerator.TokenGenerator

	logger *zerolog.Logger
}

func (c *clientService) SignUp(ctx context.Context, client types.Client) (token string, refresh string, err error) {
	client.ID = uuid.NewMD5(uuid.NameSpaceX500, []byte(client.GameID.String()+client.Email))

	client.Password = c.token.PasswordHash(client.Password)

	game, err := c.gameStorage.FindByID(client.GameID)
	if err != nil {
		return
	}

	game.CountPlayers += 1
	err = c.gameStorage.Update(game)
	if err != nil {
		return
	}

	err = c.clientStorage.CreateClient(client)
	if err != nil {
		return
	}

	token, err = c.token.GenerateToken(tokentype.Info{
		UserID: client.ID,
		GameID: client.GameID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	refresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: client.ID,
		GameID: client.GameID,
		TTL:    time.Hour * 24,
	}, tokengenerator.JWT)

	if err != nil {
		return
	}

	return token, refresh, nil
}

func (c *clientService) SignIn(ctx context.Context, client types.Client) (token, refresh string, err error) {
	client.Password = c.token.PasswordHash(client.Password)

	if client.Nickname != "" {
		client, err = c.clientStorage.FindByNicknameAndPassword(client.GameID, client.Nickname, client.Password)
		if err != nil {
			return
		}
	} else {
		client, err = c.clientStorage.FindByEmailAndPassword(client.GameID, client.Email, client.Password)
		if err != nil {
			return
		}
	}

	token, err = c.token.GenerateToken(tokentype.Info{
		UserID: client.ID,
		GameID: client.GameID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	refresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: client.ID,
		GameID: client.GameID,
		TTL:    time.Hour * 24 * 30,
	}, tokengenerator.JWT)

	if err != nil {
		return
	}

	return token, refresh, nil
}

func (c *clientService) RefreshToken(ctx context.Context, token string, refresh string) (newToken, newRefresh string, err error) {
	oldInfo, _ := c.token.ParseToken(token)

	info, err := c.token.ParseToken(refresh)
	if err != nil {
		return
	}

	client, err := c.clientStorage.FindByID(info.UserID, info.GameID)
	if err != nil {
		return
	}

	var roomID uuid.UUID
	if oldInfo.RoomID != info.RoomID {
		room, _ := c.roomStorage.FindByID(oldInfo.RoomID)
		if room.IsExists {
			roomID = room.ID
		}
	}

	newToken, err = c.token.GenerateToken(tokentype.Info{
		UserID: client.ID,
		RoomID: roomID,
		GameID: client.GameID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	newRefresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: client.ID,
		GameID: client.GameID,
		TTL:    time.Hour * 24 * 30,
	}, tokengenerator.JWT)

	return newToken, newRefresh, nil
}

func (c *clientService) GetClient(ctx context.Context, token string, gameID uuid.UUID) (client types.Client, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return types.Client{}, err
	}

	client, err = c.clientStorage.FindByID(info.UserID, gameID)
	if err != nil {
		return
	}

	client.Password = ""

	return client, err
}

func (c *clientService) UpdateClient(ctx context.Context, token string, client types.Client) (err error) {
	//info, err := c.token.ParseToken(token)
	//if err != nil {
	//	return err
	//}
	//
	//if client.Password != "" && client.Nickname != "" {
	//	oldClient, err := c.clientStorage.FindByID(client.ID, client.GameID)
	//	if err != nil {
	//		return err
	//	}
	//	oldClient.Email = client.Email
	//	if oldClient.Nickname != client.Nickname {
	//		newID := uuid.NewMD5(uuid.NameSpaceX500, []byte(client.GameID.String()+client.Email))
	//		_, err = c.clientStorage.FindByID(newID, client.GameID)
	//		if err == nil {
	//			return errors.ErrClientNickNameNotExists
	//		}
	//		oldClient.ID = newID
	//	}
	//
	//	err = c.clientStorage.Update(oldClient)
	//}
	//
	//client.ID = info.UserID
	//
	//err = c.clientStorage.Update(client)
	//if err != nil {
	//	return err
	//}

	return nil
}

func NewClientService(clientStorage storage.ClientStorage, gameStorage storage.GameStorage, roomStorage storage.RoomsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) *clientService {
	return &clientService{clientStorage: clientStorage, gameStorage: gameStorage, roomStorage: roomStorage, token: token, logger: logger}
}
