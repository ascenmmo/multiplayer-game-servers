package main

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/multiplayer-game-servers/internal/start"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := zerolog.Logger{}

	ctx := context.Background()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)
	go exiter()

	if env.DebugLogs {
		prof()
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

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

func prof() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
}

func exiter() {
	for range time.NewTicker(time.Minute * 5).C {
		os.Exit(0)
	}
}
