// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/transport/viewer"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

type loggerDevToolsGameConfigs struct {
	next multiplayer.DevToolsGameConfigs
}

func loggerMiddlewareDevToolsGameConfigs() MiddlewareDevToolsGameConfigs {
	return func(next multiplayer.DevToolsGameConfigs) multiplayer.DevToolsGameConfigs {
		return &loggerDevToolsGameConfigs{next: next}
	}
}

func (m loggerDevToolsGameConfigs) CreateOrUpdateConfig(ctx context.Context, token string, configs types.GameConfigs) (err error) {
	logger := log.Ctx(ctx).With().Str("service", "DevToolsGameConfigs").Str("method", "createOrUpdateConfig").Logger()
	defer func(begin time.Time) {
		logHandle := func(ev *zerolog.Event) {
			fields := map[string]interface{}{
				"request": viewer.Sprintf("%+v", requestDevToolsGameConfigsCreateOrUpdateConfig{
					Configs: configs,
					Token:   token,
				}),
				"response": viewer.Sprintf("%+v", responseDevToolsGameConfigsCreateOrUpdateConfig{}),
			}
			ev.Fields(fields).Str("took", time.Since(begin).String())
		}
		if err != nil {
			logger.Error().Err(err).Func(logHandle).Msg("call createOrUpdateConfig")
			return
		}
		logger.Info().Func(logHandle).Msg("call createOrUpdateConfig")
	}(time.Now())
	return m.next.CreateOrUpdateConfig(ctx, token, configs)
}

func (m loggerDevToolsGameConfigs) GetGameConfig(ctx context.Context, token string, gameID uuid.UUID) (configs types.GameConfigs, err error) {
	logger := log.Ctx(ctx).With().Str("service", "DevToolsGameConfigs").Str("method", "getGameConfig").Logger()
	defer func(begin time.Time) {
		logHandle := func(ev *zerolog.Event) {
			fields := map[string]interface{}{
				"request": viewer.Sprintf("%+v", requestDevToolsGameConfigsGetGameConfig{
					GameID: gameID,
					Token:  token,
				}),
				"response": viewer.Sprintf("%+v", responseDevToolsGameConfigsGetGameConfig{Configs: configs}),
			}
			ev.Fields(fields).Str("took", time.Since(begin).String())
		}
		if err != nil {
			logger.Error().Err(err).Func(logHandle).Msg("call getGameConfig")
			return
		}
		logger.Info().Func(logHandle).Msg("call getGameConfig")
	}(time.Now())
	return m.next.GetGameConfig(ctx, token, gameID)
}

func (m loggerDevToolsGameConfigs) GetGameResultConfigPreview(ctx context.Context, token string, gameID uuid.UUID) (gameResult types.GameConfigResults, err error) {
	logger := log.Ctx(ctx).With().Str("service", "DevToolsGameConfigs").Str("method", "getGameResultConfigPreview").Logger()
	defer func(begin time.Time) {
		logHandle := func(ev *zerolog.Event) {
			fields := map[string]interface{}{
				"request": viewer.Sprintf("%+v", requestDevToolsGameConfigsGetGameResultConfigPreview{
					GameID: gameID,
					Token:  token,
				}),
				"response": viewer.Sprintf("%+v", responseDevToolsGameConfigsGetGameResultConfigPreview{GameResult: gameResult}),
			}
			ev.Fields(fields).Str("took", time.Since(begin).String())
		}
		if err != nil {
			logger.Error().Err(err).Func(logHandle).Msg("call getGameResultConfigPreview")
			return
		}
		logger.Info().Func(logHandle).Msg("call getGameResultConfigPreview")
	}(time.Now())
	return m.next.GetGameResultConfigPreview(ctx, token, gameID)
}
