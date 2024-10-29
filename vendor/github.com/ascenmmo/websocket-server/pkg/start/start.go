package start

import (
	"context"
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/ascenmmo/websocket-server/internal/handler/tcp"
	"github.com/ascenmmo/websocket-server/internal/handler/ws"
	"github.com/ascenmmo/websocket-server/internal/service"
	configsService "github.com/ascenmmo/websocket-server/internal/service/configs_service"
	"github.com/ascenmmo/websocket-server/internal/storage"
	"github.com/ascenmmo/websocket-server/internal/utils"
	"github.com/ascenmmo/websocket-server/pkg/transport"
	"github.com/rs/zerolog"
	"time"
)

func StartWebSocket(ctx context.Context, address, tcpPort, wsPort string, token string, ratelimit int, dataTTL, gameConfigResultsTTl time.Duration, logger zerolog.Logger) (err error) {
	ramDB := memoryDB.NewMemoryDb(ctx, dataTTL)
	gameConfigResultsDB := memoryDB.NewMemoryDb(ctx, gameConfigResultsTTl)
	rateLimitDB := memoryDB.NewMemoryDb(ctx, 1)
	rateLimitBadMessageDB := memoryDB.NewMemoryDb(ctx, 1)

	tokenGenerator, err := tokengenerator.NewTokenGenerator(token)
	if err != nil {
		return err
	}

	gameConfigService := configsService.NewGameConfigsService(gameConfigResultsDB, tokenGenerator)
	newService := service.NewService(tokenGenerator, ramDB, gameConfigService, logger)

	errors := make(chan error)

	go func() {
		logger.Info().Msg(fmt.Sprintf("ws server listening on %s:%s ", address, wsPort))
		newWS := wsconnection.NewWebSocket(newService, utils.NewRateLimit(ratelimit, rateLimitDB), utils.NewRateLimit(ratelimit, rateLimitBadMessageDB), logger)
		err = newWS.Run(":" + wsPort)
		if err != nil {
			errors <- err
		}
		logger.Error().Msg("closed ws server")
	}()

	go func() {
		serverSettings := tcp.NewServerSettings(utils.NewRateLimit(ratelimit, rateLimitDB), newService)

		services := []transport.Option{
			transport.MaxBodySize(10 * 1024 * 1024),
			transport.ServerSettings(transport.NewServerSettings(serverSettings)),
		}

		srv := transport.New(logger, services...).WithLog()
		logger.Info().Msg(fmt.Sprintf("rest ws server listening on %s:%s ", address, tcpPort))
		if err := srv.Fiber().Listen(":" + tcpPort); err != nil {
			errors <- err
		}
	}()
	err = <-errors

	return err
}
