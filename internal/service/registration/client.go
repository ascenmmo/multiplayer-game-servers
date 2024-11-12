package registration

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type clientService struct {
	clientStorage    storage.ClientStorage
	gameStorage      storage.GameStorage
	roomStorage      storage.RoomsStorage
	gameSavesStorage storage.GameSavesStorage
	token            tokengenerator.TokenGenerator

	logger *zerolog.Logger
}

func (c *clientService) SignUp(ctx context.Context, client types.Client) (token string, refresh string, err error) {
	newClient := types.Client{}
	newClient.Update(client)

	newClient.GameID = client.GameID
	newClient.Password = client.Password
	newClient.NewPassword = client.NewPassword

	newClient.Password, err = c.token.GenerateSecretHash(newClient.Password)
	if err != nil {
		return token, refresh, errors.ErrBadNewPassword
	}

	game, err := c.gameStorage.FindByID(newClient.GameID)
	if err != nil {
		return token, refresh, errors.ErrGameNotFound
	}

	game.CountPlayers += 1
	err = c.gameStorage.Update(game)
	if err != nil {
		return token, refresh, errors.ErrGameNotFound
	}

	err = c.clientStorage.CreateClient(newClient)
	if err != nil {
		return token, refresh, errors.ErrClientCreationError
	}

	token, err = c.token.GenerateToken(tokentype.Info{
		UserID: newClient.ID,
		GameID: newClient.GameID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	refresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: newClient.ID,
		GameID: newClient.GameID,
		TTL:    time.Hour * 24,
	}, tokengenerator.JWT)

	if err != nil {
		return
	}

	return token, refresh, nil
}

func (c *clientService) SignIn(ctx context.Context, client types.Client) (token, refresh string, err error) {
	client.Password, err = c.token.GenerateSecretHash(client.Password)
	if err != nil {
		return token, refresh, errors.ErrBadNewPassword
	}

	client, err = c.clientStorage.FindByPassword(client.GameID, client.Email, client.Nickname, client.Password)
	if err != nil {
		return token, refresh, errors.ErrWrongUserOrPassword
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
	info, err := c.token.ParseToken(token)
	if err != nil {
		return err
	}

	oldClient, err := c.clientStorage.FindByID(info.UserID, info.GameID)
	if err != nil {
		return errors.ErrNotFound
	}

	if client.Password != "" && client.NewPassword != "" {
		hash, _ := c.token.GenerateSecretHash(client.Password)
		newPswHash, err := c.token.GenerateSecretHash(client.NewPassword)
		if err != nil {
			return errors.ErrBadNewPassword
		}
		if oldClient.Password != hash {
			return errors.ErrNotFound
		}
		oldClient.Password = newPswHash
	}

	oldClient.Update(client)

	client.ID = info.UserID

	err = c.clientStorage.Update(client)
	if err != nil {
		return err
	}

	return nil
}

func (c *clientService) GetGameSaves(ctx context.Context, token string) (saves types.GameSaves, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return saves, err
	}
	saves, err = c.gameSavesStorage.FindByID(info.GameID, info.UserID)
	if err != nil {
		return saves, errors.ErrGameSaves
	}

	return saves, nil
}

func (c *clientService) SetGameSaves(ctx context.Context, token string, saves types.GameSaves) (err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return err
	}

	saves.ID = uuid.NewMD5(uuid.NameSpaceX500, []byte(info.GameID.String()+info.UserID.String()))
	saves.GameID = info.GameID
	saves.UserID = info.UserID

	_, err = c.gameSavesStorage.FindByID(info.GameID, info.UserID)
	if err != nil {
		return c.gameSavesStorage.Create(saves)
	}

	return c.gameSavesStorage.Update(saves)
}

func (c *clientService) DeleteGameSaves(ctx context.Context, token string) (err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return err
	}

	return c.gameSavesStorage.Delete(info.GameID, info.UserID)
}

func NewClientService(clientStorage storage.ClientStorage, gameStorage storage.GameStorage, gameSavesStorage storage.GameSavesStorage, roomStorage storage.RoomsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) *clientService {
	return &clientService{clientStorage: clientStorage, gameStorage: gameStorage, gameSavesStorage: gameSavesStorage, roomStorage: roomStorage, token: token, logger: logger}
}
