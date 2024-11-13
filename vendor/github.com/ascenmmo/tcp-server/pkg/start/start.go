package start

import (
	"context"
	"fmt"
	"github.com/ascenmmo/tcp-server/internal/handler"
	"github.com/ascenmmo/tcp-server/internal/service"
	"github.com/ascenmmo/tcp-server/internal/storage"
	"github.com/ascenmmo/tcp-server/internal/utils"
	"github.com/ascenmmo/tcp-server/pkg/transport"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/rs/zerolog"
	"runtime"
	"time"
)

func StartTCP(ctx context.Context, address string, port string, token string, rateLimit int, dataTTL time.Duration, logger zerolog.Logger, logWithMemoryUsage bool) (err error) {
	ramDB := memoryDB.NewMemoryDb(ctx, dataTTL)
	rateLimitDB := memoryDB.NewMemoryDb(ctx, 1)

	tokenGenerator, err := tokengenerator.NewTokenGenerator(token)
	if err != nil {
		return err
	}

	newService := service.NewService(tokenGenerator, ramDB, logger)

	rateLimitSerice := utils.NewRateLimit(rateLimit, rateLimitDB)

	connection := handler.NewRestConnection(rateLimitSerice, newService)
	serverSettings := handler.NewServerSettings(rateLimitSerice, newService)

	if logWithMemoryUsage {
		logMemoryUsage(logger)
	}

	services := []transport.Option{
		transport.MaxBodySize(10 * 1024 * 1024),
		transport.GameConnections(transport.NewGameConnections(connection)),
		transport.ServerSettings(transport.NewServerSettings(serverSettings)),
	}

	srv := transport.New(logger, services...).WithLog()

	logger.Info().Msg(fmt.Sprintf("rest game server listening on %s:%s ", address, port))

	return srv.Fiber().Listen(":" + port)
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
