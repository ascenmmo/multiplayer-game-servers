package start

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/websocket-server/pkg/start"
	"github.com/rs/zerolog"
)

func WebsocketRun(ctx context.Context, logger zerolog.Logger) {
	err := start.StartWebSocket(
		ctx,
		env.ServerAddress,
		env.WebsocketServerPort,
		env.WebsocketServerConnectionPort,
		env.TokenKey,
		env.WebsocketServerMaxRequestPerSecond,
		10,
		logger.With().Str("server:", "websocket").Logger(),
		false,
	)
	mastNil(err)
}
