package start

import (
	"context"
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/ascenmmo/udp-server/internal/handler/tcp"
	"github.com/ascenmmo/udp-server/internal/handler/udp"
	"github.com/ascenmmo/udp-server/internal/service"
	configsService "github.com/ascenmmo/udp-server/internal/service/configs_service"
	memoryDB "github.com/ascenmmo/udp-server/internal/storage"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/ascenmmo/udp-server/pkg/transport"
	"github.com/rs/zerolog"
	"time"
)

func StartUDP(ctx context.Context, address string, tcpPort, udpPort string, token string, udpRateLimit int, dataTTL, gameConfigResultsTTl time.Duration, logger zerolog.Logger) (err error) {
	ramDB := memoryDB.NewMemoryDb(ctx, dataTTL)
	gameConfigResultsDB := memoryDB.NewMemoryDb(ctx, gameConfigResultsTTl)
	rateLimitDB := memoryDB.NewMemoryDb(ctx, 1)

	tokenGen, err := tokengenerator.NewTokenGenerator(token)
	if err != nil {
		return err
	}

	gameConfigsService := configsService.NewGameConfigsService(gameConfigResultsDB, tokenGen)
	newService := service.NewService(tokenGen, ramDB, gameConfigsService, logger)

	errors := make(chan error)

	newUDP, err := udp.NewWorkerUDP(udpPort, newService, udpRateLimit, rateLimitDB, logger)
	if err != nil {
		return err
	}

	go func() {
		err := newUDP.Listener(ctx)
		if err != nil {
			errors <- err
		}
		logger.Error().Msg("closed udp listen server")
	}()

	go newUDP.Sender(ctx)

	go func() {
		serverSettings := tcp.NewServerSettings(utils.NewRateLimit(10, rateLimitDB), newService)

		services := []transport.Option{
			transport.MaxBodySize(10 * 1024 * 1024),
			transport.ServerSettings(transport.NewServerSettings(serverSettings)),
		}

		srv := transport.New(logger, services...).WithLog()

		logger.Info().Msg(fmt.Sprintf("udp rest server listening on %s:%s ", address, tcpPort))
		if err := srv.Fiber().Listen(":" + tcpPort); err != nil {
			errors <- err
		}
	}()

	err = <-errors

	return err
}
