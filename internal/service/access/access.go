package access

import (
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type AccessGame interface {
	CreateNew(gameID, userID uuid.UUID) (err error)
	AddOwner(ownerID uuid.UUID, gameID, newOwnerID uuid.UUID) (err error)
	GetOwnerAccess(gameID, ownerID uuid.UUID) (err error)
	GetOwnerGames(ownerID uuid.UUID) (gameIDs []uuid.UUID, err error)
	RemoveUser(gameID, ownerID, userID uuid.UUID) (err error)
}

type accessGame struct {
	storage storage.AccessGameStorage
}

func (c *accessGame) CreateNew(gameID, userID uuid.UUID) (err error) {
	gameAccess := types.AccessGame{
		ID:        uuid.NewMD5(uuid.New(), []byte(gameID.String())),
		GameID:    gameID,
		CreatorID: userID,
		Owners:    []uuid.UUID{userID},
	}

	err = c.storage.CreateAccessGame(gameAccess)
	if err != nil {
		return err
	}

	return nil
}

func (c *accessGame) AddOwner(ownerID uuid.UUID, gameID, newOwnerID uuid.UUID) (err error) {
	gameAccess, err := c.storage.FindByGameID(gameID)
	if err != nil {
		return err
	}
	isOwner := false
	for _, owner := range gameAccess.Owners {
		if ownerID == owner {
			isOwner = true
		}
	}

	if !isOwner {
		return errors.ErrAccessDenied
	}

	gameAccess.Owners = append(gameAccess.Owners, newOwnerID)
	err = c.storage.CreateAccessGame(gameAccess)
	if err != nil {
		return err
	}

	return nil
}

func (c *accessGame) GetOwnerAccess(gameID, ownerID uuid.UUID) (err error) {
	gameAccess, err := c.storage.FindByGameID(gameID)
	if err != nil {
		return err
	}

	isOwner := false
	for _, owner := range gameAccess.Owners {
		if ownerID == owner {
			isOwner = true
		}
	}

	if !isOwner {
		return errors.ErrAccessDenied
	}

	return nil
}

func (c *accessGame) GetOwnerGames(ownerID uuid.UUID) (gameIDs []uuid.UUID, err error) {
	games, err := c.storage.FindByOwnerID(ownerID)
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, errors.ErrAccessDenied
	}

	for _, v := range games {
		gameIDs = append(gameIDs, v.GameID)
	}

	return gameIDs, nil
}

func (c *accessGame) RemoveUser(gameID, ownerID, userID uuid.UUID) (err error) {
	gameAccess, err := c.storage.FindByGameID(gameID)
	if err != nil {
		return err
	}

	isOwner := false
	for _, owner := range gameAccess.Owners {
		if ownerID == owner {
			isOwner = true
		}
	}

	if !isOwner {
		return errors.ErrAccessDenied
	}

	if userID == gameAccess.CreatorID {
		return errors.ErrAccessDeniedDeleteCreatorID
	}

	gameAccess.RemoveOwner(userID)

	err = c.storage.Update(gameAccess)
	if err != nil {
		return err
	}

	return nil
}

func NewAccessGame(storage storage.AccessGameStorage) *accessGame {
	return &accessGame{storage: storage}
}
