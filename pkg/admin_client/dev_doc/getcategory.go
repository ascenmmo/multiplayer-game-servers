package dev_doc

import (
	ruclient "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/client"
	ruconnections "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/connections"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/tcp"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/udp"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/ru/ws"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetCategory() []types.DocCategory {
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

	connectionService := udp.GetUPDConnectionsStruct()
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
