package devtools

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/service/access"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type gameConfigs struct {
	accessGame         access.AccessGame
	gameConfigsStorage storage.GameConfigsStorage

	token tokengenerator.TokenGenerator

	logger *zerolog.Logger
}

func (g *gameConfigs) CreateOrUpdateConfig(ctx context.Context, token string, configs types.GameConfigs) (err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return err
	}

	err = g.accessGame.GetOwnerAccess(configs.GameID, info.UserID)
	if err != nil {
		return err
	}

	err = g.gameConfigsStorage.CreateOrUpdateNewConfig(configs)
	if err != nil {
		return err
	}

	return nil
}

func (g *gameConfigs) GetGameConfig(ctx context.Context, token string, gameID uuid.UUID) (configs types.GameConfigs, err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return configs, err
	}

	err = g.accessGame.GetOwnerAccess(info.GameID, info.UserID)
	if err != nil {
		return configs, err
	}

	configs, err = g.gameConfigsStorage.GetConfig(gameID)
	if err != nil {
		return configs, err
	}

	return configs, nil
}

func (g *gameConfigs) schedulerConfigRunner() {
	//tiker := time.NewTicker(time.Second * 10)

}

func NewGameConfigs(accessGame access.AccessGame, gameConfigsStorage storage.GameConfigsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.DevToolsGameConfigs {
	return &gameConfigs{accessGame: accessGame, gameConfigsStorage: gameConfigsStorage, token: token, logger: logger}
}
