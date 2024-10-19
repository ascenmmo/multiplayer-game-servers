package start

import (
	"context"
	"fmt"
	"github.com/ascenmmo/tcp-server/internal/handler"
	"github.com/ascenmmo/tcp-server/internal/service"
	configsService "github.com/ascenmmo/tcp-server/internal/service/configs_service"
	"github.com/ascenmmo/tcp-server/internal/storage"
	"github.com/ascenmmo/tcp-server/internal/utils"
	"github.com/ascenmmo/tcp-server/pkg/transport"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/rs/zerolog"
	"time"
)

func StartTCP(ctx context.Context, address string, port string, token string, rateLimit int, dataTTL, gameConfigResultsTTl time.Duration, logger zerolog.Logger) (err error) {
	ramDB := memoryDB.NewMemoryDb(ctx, dataTTL)
	gameConfigResultsDB := memoryDB.NewMemoryDb(ctx, gameConfigResultsTTl)
	rateLimitDB := memoryDB.NewMemoryDb(ctx, 1)

	tokenGenerator, err := tokengenerator.NewTokenGenerator(token)
	if err != nil {
		return err
	}

	gameConfigsService := configsService.NewGameConfigsService(gameConfigResultsDB, tokenGenerator)
	newService := service.NewService(tokenGenerator, ramDB, gameConfigsService, logger)

	rateLimitSerice := utils.NewRateLimit(rateLimit, rateLimitDB)

	connection := handler.NewRestConnection(rateLimitSerice, newService)
	serverSettings := handler.NewServerSettings(rateLimitSerice, newService)

	services := []transport.Option{
		transport.MaxBodySize(10 * 1024 * 1024),
		transport.GameConnections(transport.NewGameConnections(connection)),
		transport.ServerSettings(transport.NewServerSettings(serverSettings)),
	}

	srv := transport.New(logger, services...).WithLog()

	logger.Info().Msg(fmt.Sprintf("rest game server listening on %s:%s ", address, port))

	return srv.Fiber().Listen(":" + port)
}
