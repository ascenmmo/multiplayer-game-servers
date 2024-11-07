package start

import (
	"context"
	"fmt"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/ascenmmo/udp-server/internal/handler/tcp"
	"github.com/ascenmmo/udp-server/internal/handler/udp"
	"github.com/ascenmmo/udp-server/internal/service"
	memoryDB "github.com/ascenmmo/udp-server/internal/storage"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/ascenmmo/udp-server/pkg/transport"
	"github.com/rs/zerolog"
	"runtime"
	"time"
)

func StartUDP(ctx context.Context, address string, tcpPort, udpPort string, token string, udpRateLimit int, dataTTL time.Duration, logger zerolog.Logger, logWithMemoryUsage bool) (err error) {
	ramDB := memoryDB.NewMemoryDb(ctx, dataTTL)
	rateLimitDB := memoryDB.NewMemoryDb(ctx, 1)

	tokenGen, err := tokengenerator.NewTokenGenerator(token)
	if err != nil {
		return err
	}

	newService := service.NewService(tokenGen, ramDB, logger)

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

	if logWithMemoryUsage {
		logMemoryUsage(logger)
	}

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
