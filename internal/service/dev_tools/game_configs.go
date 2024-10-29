package devtools

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	"github.com/ascenmmo/multiplayer-game-servers/internal/service/access"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"strings"
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

	for i, conf := range configs.SortingConfig {
		for i := i + 1; i < len(configs.SortingConfig); i++ {
			if conf.ResultName == configs.SortingConfig[i].ResultName {
				return errors.ErrGameConfigSameResultName
			}
		}
		for j, cloumn := range conf.Params {
			for j := j + 1; j < len(conf.Params); j++ {
				if cloumn.ColumnName == conf.Params[j].ColumnName {
					return errors.ErrGameConfigSameColumnName
				}
			}
		}
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

	err = g.accessGame.GetOwnerAccess(gameID, info.UserID)
	if err != nil {
		return configs, err
	}

	configs, err = g.gameConfigsStorage.GetConfig(gameID)
	if err != nil {
		return configs, err
	}

	return configs, nil
}

func (g *gameConfigs) GetGameResultConfigPreview(ctx context.Context, token string, gameID uuid.UUID) (gameResult types.GameConfigResults, err error) {
	info, err := g.token.ParseToken(token)
	if err != nil {
		return gameResult, err
	}

	err = g.accessGame.GetOwnerAccess(gameID, info.UserID)
	if err != nil {
		return gameResult, err
	}

	configs, err := g.gameConfigsStorage.GetConfig(gameID)
	if err != nil {
		return gameResult, err
	}

	gameResult.GameID = gameID
	gameResult.RoomID = uuid.New()

	gameResult.Result = make(map[string]interface{}, len(configs.SortingConfig))

	for _, v := range configs.SortingConfig {
		d := strings.ToLower(v.ResultType)
		switch d {
		case "string":
			gameResult.Result[v.ResultName] = "someData"
		case "int":
			gameResult.Result[v.ResultName] = 12345
		case "float":
			gameResult.Result[v.ResultName] = 12345.54321
		}
	}

	return gameResult, nil
}

func (g *gameConfigs) schedulerConfigRunner() {
	//tiker := time.NewTicker(time.Second * 10)

}

func NewGameConfigs(accessGame access.AccessGame, gameConfigsStorage storage.GameConfigsStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.DevToolsGameConfigs {
	return &gameConfigs{accessGame: accessGame, gameConfigsStorage: gameConfigsStorage, token: token, logger: logger}
}
