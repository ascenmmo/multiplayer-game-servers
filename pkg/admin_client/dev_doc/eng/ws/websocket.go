package engws

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetWsConnectionsStruct() []types.DocStruct {
	ws := types.DocStruct{
		Title: "WebSocket Connection",
		Info: "This WebSocket service is specifically designed for multiplayer online games, enabling real-time data exchange between clients. " +
			"It facilitates quick and efficient sharing of game data, which is crucial in scenarios where speed takes precedence over message delivery reliability. " +
			"WebSocket supports bidirectional communication between the client and the server, making it an excellent choice for game clients. However, packet loss or out-of-order delivery may occur.",
		DockLists: []types.DocStructList{
			{
				Title: "Sending and Receiving Messages",
				Description: "To connect to the game WebSocket service, you need to know the URL and use a token for authentication. " +
					"The maximum number of requests per second is 200. " +
					"Example command for connecting and sending a game message via WebSocket:",
				RequestPath:   "ws://127.0.0.1:4240/api/ws/connect",
				RequestHeader: "Token: your token",
				Method:        "ws",
				RequestBody: `{
	"data": "data to be sent"
}`,
				RequestBodyInfo: "The request body should contain game information to be sent to other players. The authentication token is passed in the header.",
				ResponseBody: `{
	"data": "data from other clients"
}`,
				ResponseBodyInfo: "The response contains data received from other game clients connected to the same game room.",
			},
		},
	}

	return []types.DocStruct{ws}
}
