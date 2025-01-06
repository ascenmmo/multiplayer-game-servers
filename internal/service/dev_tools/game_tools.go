package devtools

import (
	"context"
	"fmt"
	"github.com/ascenmmo/multiplayer-game-servers/internal/service/access"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type devTools struct {
	accessGame       access.AccessGame
	gameStorage      storage.GameStorage
	serverStorage    storage.ServersStorage
	roomStorage      storage.RoomsStorage
	developerStorage storage.DeveloperStorage

	token tokengenerator.TokenGenerator

	logger *zerolog.Logger
}

func (g *devTools) CreateGame(ctx context.Context, token string, game types.Game) (id uuid.UUID, err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return id, err
	}

	game.ID = uuid.NewMD5(uuid.NameSpaceX500, []byte(game.Name))
	game.Owners = append(game.Owners, info.UserID)

	err = g.gameStorage.CreateGame(game)
	if err != nil {
		g.accessGame.RemoveUser(game.ID, info.UserID, info.UserID)
		return id, err
	}
	game.RemoveEmptyArray()

	err = g.accessGame.CreateNew(game.ID, info.UserID)
	if err != nil {
		return id, err
	}

	return game.ID, err
}

func (g *devTools) GameAddOwner(ctx context.Context, token string, gameID, userID uuid.UUID) (err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return err
	}

	game, err := g.gameStorage.FindByID(gameID)
	if err != nil {
		return err
	}

	game.Owners = append(game.Owners, userID)

	err = g.gameStorage.Update(game)
	if err != nil {
		return err
	}

	err = g.accessGame.AddOwner(info.UserID, gameID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (g *devTools) GameRemoveOwner(ctx context.Context, token string, gameID, userID uuid.UUID) (err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return err
	}

	err = g.accessGame.RemoveUser(gameID, info.UserID, userID)
	if err != nil {
		return err
	}

	game, err := g.gameStorage.FindByID(gameID)
	if err != nil {
		return err
	}

	game.RemoveOwners(userID)

	err = g.gameStorage.Update(game)
	if err != nil {
		return err
	}

	return nil
}

func (g *devTools) UpdateGame(ctx context.Context, token string, gameID uuid.UUID, newGame types.Game) (id uuid.UUID, err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return id, err
	}

	err = g.accessGame.GetOwnerAccess(gameID, info.UserID)
	if err != nil {
		return id, err
	}

	game, err := g.gameStorage.FindByID(gameID)
	if err != nil {
		return id, err
	}

	if !game.UpdateGame(newGame) {
		return gameID, err
	}

	err = g.gameStorage.Update(game)
	if err != nil {
		return id, err
	}

	return game.ID, nil
}

func (g *devTools) DeleteGame(ctx context.Context, token string, gameID uuid.UUID) (err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return err
	}

	err = g.accessGame.RemoveGame(gameID, info.UserID)
	if err != nil {
		return err
	}

	err = g.gameStorage.DeleteGame(gameID)
	if err != nil {
		return err
	}
	return nil
}

func (g *devTools) GetMyGames(ctx context.Context, token string) (games []types.Game, err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return games, err
	}

	games, err = g.gameStorage.FindByOwnerID(info.UserID)
	if err != nil {
		return games, err
	}
	for i := range games {
		games[i].RemoveEmptyArray()
	}

	return games, nil
}

func (g *devTools) GetGameByGameID(ctx context.Context, token string, gameID uuid.UUID) (game types.Game, err error) {
	_, err = g.token.ParseToken(token)
	if err != nil {
		return game, err
	}

	game, err = g.gameStorage.FindByID(gameID)
	if err != nil {
		return game, err
	}

	rooms, err := g.roomStorage.FindAll(game.ID)
	if err != nil {
		g.logger.Error().Err(err).Msg("err get rooms GetGameByGameID")
	}

	var countOnline int
	for _, v := range rooms {
		countOnline += len(v.Connections)
	}
	game.CountOnline = countOnline
	err = g.gameStorage.Update(game)
	if err != nil {
		g.logger.Error().Err(err).Msg("err update game GetGameByGameID")
	}

	game.RemoveEmptyArray()

	developers, err := g.developerStorage.FindByIDS(game.Owners)
	if err != nil {
		return game, nil
	}

	for _, v := range developers {
		txt := fmt.Sprintf(`"%s":"%s"`, v.Nickname, v.ID)
		game.OwnersInfoLint = append(game.OwnersInfoLint, txt)
	}

	return
}

func (g *devTools) TurnOnServerInGame(ctx context.Context, token string, serverID uuid.UUID, gameId uuid.UUID) (err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return err
	}

	err = g.accessGame.GetOwnerAccess(gameId, info.UserID)
	if err != nil {
		return err
	}

	game, err := g.gameStorage.FindByID(gameId)
	if err != nil {
		return err
	}

	game.TurnOnServer(serverID)

	err = g.gameStorage.Update(game)

	return nil
}

func (g *devTools) TurnOffServerInGame(ctx context.Context, token string, serverID uuid.UUID, gameId uuid.UUID) (err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return err
	}

	err = g.accessGame.GetOwnerAccess(gameId, info.UserID)
	if err != nil {
		return err
	}

	game, err := g.gameStorage.FindByID(gameId)
	if err != nil {
		return err
	}

	game.RemoveServer(serverID)

	err = g.gameStorage.Update(game)

	return nil
}

func NewDevTools(accessGame access.AccessGame, gameStorage storage.GameStorage, serverStorage storage.ServersStorage, roomStorage storage.RoomsStorage, developerStorage storage.DeveloperStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.DevTools {
	return &devTools{accessGame: accessGame, gameStorage: gameStorage, serverStorage: serverStorage, developerStorage: developerStorage, roomStorage: roomStorage, token: token, logger: logger}
}
