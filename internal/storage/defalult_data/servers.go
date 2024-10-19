package defalultdata

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

var (
	defaultServerUDP = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerUDP")),
		ServerName:     "ascenmmo udp",
		Address:        "ascenmmo.com:4500",
		ServerType:     types.ServerTypeUDP,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}

	defaultServerWebsocket = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerWebsocket")),
		ServerName:     "ascenmmo ws",
		Address:        "ascenmmo.com:4240",
		ServerType:     types.ServerTypeWebsocket,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}

	defaultServerTCP = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerTCP")),
		ServerName:     "ascenmmo tcp",
		Address:        "ascenmmo.com:8080",
		ServerType:     types.ServerTypeTCP,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}

	defaultServerChatWs = types.Server{
		ID:             uuid.NewMD5(uuid.UUID{}, []byte("defaultServerTCP")),
		ServerName:     "ascenmmo tcp",
		Address:        "ascenmmo.com:4244",
		ServerType:     types.ServerTypeChat,
		Region:         "europe",
		IsActive:       true,
		AscenmmoServer: true,
	}
)

func AddServers(creatorID uuid.UUID) (servers []types.Server) {
	defaultServers := []types.Server{
		defaultServerUDP,
		defaultServerWebsocket,
		defaultServerTCP,
		defaultServerChatWs,
	}
	for _, server := range defaultServers {
		server.Owners = []uuid.UUID{creatorID}
		servers = append(servers, server)
	}

	return servers
}
