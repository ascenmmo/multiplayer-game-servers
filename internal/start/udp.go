package start

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/udp-server/pkg/start"
	"github.com/rs/zerolog"
)

func UdpServer(ctx context.Context, logger zerolog.Logger) {
	err := start.StartUDP(
		ctx,
		env.ServerAddress,
		env.UdpServerPort,
		env.UdpServerConnectionPort,
		env.TokenKey,
		env.UdpServerMaxRequestPerSecond,
		10,
		60,
		logger.With().Str("server:", "udp").Logger(),
	)
	mastNil(err)
}
