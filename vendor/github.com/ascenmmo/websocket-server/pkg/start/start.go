package start

import (
	"context"
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/ascenmmo/websocket-server/internal/handler/tcp"
	"github.com/ascenmmo/websocket-server/internal/handler/ws"
	"github.com/ascenmmo/websocket-server/internal/service"
	"github.com/ascenmmo/websocket-server/internal/storage"
	"github.com/ascenmmo/websocket-server/internal/utils"
	"github.com/ascenmmo/websocket-server/pkg/transport"
	"github.com/rs/zerolog"
	"runtime"
	"time"
)

func StartWebSocket(ctx context.Context, address, tcpPort, wsPort string, token string, ratelimit int, dataTTL time.Duration, logger zerolog.Logger, logWithMemoryUsage bool) (err error) {
	ramDB := memoryDB.NewMemoryDb(ctx, dataTTL)
	rateLimitDB := memoryDB.NewMemoryDb(ctx, 1)

	tokenGenerator, err := tokengenerator.NewTokenGenerator(token)
	if err != nil {
		return err
	}

	newService := service.NewService(tokenGenerator, ramDB, logger)

	errors := make(chan error)

	if logWithMemoryUsage {
		logMemoryUsage(logger)
	}

	go func() {
		logger.Info().Msg(fmt.Sprintf("ws server listening on %s:%s ", address, wsPort))
		newWS := wsconnection.NewWebSocket(newService, utils.NewRateLimit(ratelimit, rateLimitDB), logger)
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

func logMemoryUsage(logger zerolog.Logger) {
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for range ticker.C {
			var stats runtime.MemStats
			runtime.ReadMemStats(&stats)
			logger.Info().
				Interface("num cpu", runtime.NumCPU()).
				Interface("Memory Usage", stats.Alloc/1024/1024).
				Interface("TotalAlloc", stats.TotalAlloc/1024/1024).
				Interface("Sys", stats.Sys/1024/1024).
				Interface("NumGC", stats.NumGC)
		}
	}()
}
