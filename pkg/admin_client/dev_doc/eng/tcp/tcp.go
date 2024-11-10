package engtcp

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetTCPConnectionsStruct() []types.DocStruct {
	tcp := types.DocStruct{
		Title: "TCP Connection",
		Info: "This service is designed for interaction between game clients and the server via a TCP connection. " +
			"It allows clients to send messages, receive new messages, and remove users. " +
			"The service supports up to **10 requests per second** and is ideal for game applications " +
			"that require high-speed interaction and long polling for retrieving new messages.",
		DockLists: []types.DocStructList{
			{
				Title:         "Sending Messages",
				Description:   "To send a message from the client to the server, use the following endpoint. You must include the token in the request header.",
				RequestPath:   "/api/v1/rest/gameConnections/setSendMessage",
				RequestHeader: "Token: your token",
				Method:        "POST",
				RequestBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "params": {
    "message": {
      "data": "your data",
      "server": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "token": "your token"
    }
  }
}`,
				RequestBodyInfo: "The request body should include message data, server identifier, and authentication token.",
				ResponseBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {}
}`,
				ResponseBodyInfo: "The response indicates successful operation.",
			},
			{
				Title:         "Receiving New Messages",
				Description:   "To receive new messages from the server, use the following endpoint. Long polling can be used.",
				RequestPath:   "/api/v1/rest/gameConnections/getMessage",
				RequestHeader: "Token: your token",
				Method:        "POST",
				RequestBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "params": {}
}`,
				RequestBodyInfo: "The request body does not require parameters, only the token in the header is needed.",
				ResponseBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {
    "messages": {
      "dataArray": [
        "new message"
      ]
    }
  }
}`,
				ResponseBodyInfo: "The response contains an array of new messages from the server.",
			},
			{
				Title:         "Removing a User",
				Description:   "To remove a user from the game server, use the following endpoint.",
				RequestPath:   "/api/v1/rest/gameConnections/removeUser",
				RequestHeader: "Token: your token",
				Method:        "POST",
				RequestBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "params": {
    "userID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
  }
}`,
				RequestBodyInfo: "The request body should include the unique identifier of the user to be removed.",
				ResponseBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {}
}`,
				ResponseBodyInfo: "The response indicates successful operation.",
			},
		},
	}

	return []types.DocStruct{tcp}
}
