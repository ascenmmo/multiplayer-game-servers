package start

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/tcp-server/pkg/start"
	"github.com/rs/zerolog"
)

func TcpServer(ctx context.Context, logger zerolog.Logger) {
	err := start.StartTCP(
		ctx,
		env.ServerAddress,
		env.TcpServerPort,
		env.TokenKey,
		env.TcpServerMaxRequestPerSecond,
		10,
		logger.With().Str("server:", "tcp").Logger(),
		false,
	)
	mastNil(err)
}
