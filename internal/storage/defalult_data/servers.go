package defalultdata

import (
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

var (
	defaultServerUDP = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerUDP")),
		ServerName:     "server udp",
		Address:        env.ServerAddress + ":" + env.UdpServerPort,
		ServerType:     types.ServerTypeUDP,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}

	defaultServerWebsocket = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerWebsocket")),
		ServerName:     "server ws",
		Address:        env.ServerAddress + ":" + env.WebsocketServerPort,
		ServerType:     types.ServerTypeWebsocket,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}

	defaultServerTCP = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerTCP")),
		ServerName:     "server tcp",
		Address:        env.ServerAddress + ":" + env.TcpServerPort,
		ServerType:     types.ServerTypeTCP,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}
)

func AddServers(creatorID uuid.UUID) (servers []types.Server) {
	defaultServers := []types.Server{}

	if env.RunUdpServer {
		defaultServers = append(defaultServers, defaultServerUDP)
	}

	if env.RunWebsocketServer {
		defaultServers = append(defaultServers, defaultServerWebsocket)
	}

	if env.RunTcpServer {
		defaultServers = append(defaultServers, defaultServerTCP)
	}

	for _, server := range defaultServers {
		server.Owners = []uuid.UUID{creatorID}
		servers = append(servers, server)
	}

	return servers
}
