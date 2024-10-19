package main

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/multiplayer-game-servers/internal/start"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	ctx := context.Background()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	if env.RunMultiplayer {
		go start.Multiplayer(logger)
	}

	if env.RunUdpServer {
		go start.UdpServer(ctx, logger)
	}

	if env.RunWebsocketServer {
		go start.WebsocketRun(ctx, logger)
	}

	if env.RunTcpServer {
		go start.TcpServer(ctx, logger)
	}

	<-shutdown
}
