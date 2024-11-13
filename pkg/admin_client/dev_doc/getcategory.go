package dev_doc

import (
	engclient "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/eng/client"
	engconnections "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/eng/connections"
	engtcp "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/eng/tcp"
	engudp "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/eng/udp"
	engws "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/eng/ws"
	ruclient "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/client"
	ruconnections "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/connections"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/tcp"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/udp"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/ws"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetCategory(lang string) []types.DocCategory {
	switch lang {
	case "ru":
		return RuGetCategory()
	case "eng":
		return EngGetCategory()
	default:
		return EngGetCategory()
	}
}

func RuGetCategory() []types.DocCategory {
	client := types.DocCategory{
		CategoryTitle: "Клиенты",
		DocStruct:     ruclient.GetClientStruct(),
		DocErrorList:  ruclient.ClientErrors(),
	}

	connection := types.DocCategory{
		CategoryTitle: "Создание подключений",
		DocStruct:     ruconnections.GetConnectionsStruct(),
		DocErrorList:  ruconnections.ConnectionErrors(),
	}

	connectionService := udp.GetUDPConnectionsStruct()
	connectionService = append(connectionService, ws.GetWsConnectionsStruct()...)
	connectionService = append(connectionService, tcp.GetTCPConnectionsStruct()...)

	connectionServiceCategory := types.DocCategory{
		CategoryTitle: "Подключение к сервисам",
		DocStruct:     connectionService,
		DocErrorList:  tcp.TcpErrors(),
	}

	data := []types.DocCategory{
		client,
		connection,
		connectionServiceCategory,
	}
	return data
}

func EngGetCategory() []types.DocCategory {
	client := types.DocCategory{
		CategoryTitle: "Clients",
		DocStruct:     engclient.GetClientStruct(),
		DocErrorList:  engclient.ClientErrors(),
	}

	connection := types.DocCategory{
		CategoryTitle: "Create new connection",
		DocStruct:     engconnections.GetConnectionsStruct(),
		DocErrorList:  engconnections.ConnectionErrors(),
	}

	connectionService := engudp.GetUDPConnectionsStruct()
	connectionService = append(connectionService, engws.GetWsConnectionsStruct()...)
	connectionService = append(connectionService, engtcp.GetTCPConnectionsStruct()...)

	connectionServiceCategory := types.DocCategory{
		CategoryTitle: "Connect to servers",
		DocStruct:     connectionService,
		DocErrorList:  engtcp.TcpErrors(),
	}

	data := []types.DocCategory{
		client,
		connection,
		connectionServiceCategory,
	}
	return data
}
