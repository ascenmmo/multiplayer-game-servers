package engconnections

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetConnectionsStruct() []types.DocStruct {
	creteRoomDoc := types.DocStruct{
		Title: "Create Room",
		Info:  "This method is used to create a new room in the game. The request requires a token in the header for authorization. The response returns a new token, which can be used for subsequent requests.",
		DockLists: []types.DocStructList{
			{
				Title:            "Create Room",
				Description:      "Request to create a new room for the game. The request must be authorized using a token.",
				RequestPath:      "/api/v1/devToolsConnections/createRoom",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"name\": \"string\"\n  }\n}",
				RequestBodyInfo:  "The request body contains the name of the room (name). You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "The response returns a new token that can be used for subsequent requests.",
			},
		},
	}

	getMyRoomDoc := types.DocStruct{
		Title: "Get Information About Your Room",
		Info:  "This method is used to get information about the current user's room. The request requires a token in the header for authorization. The response returns the room data, including information about connections, creation time, servers, and other parameters.",
		DockLists: []types.DocStructList{
			{
				Title:            "Get My Room",
				Description:      "Request to get information about the current user's room. The request must be authorized using a token.",
				RequestPath:      "/api/v1/devToolsConnections/getMyRoom",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "The request body is empty because the data is retrieved for the currently authorized user. You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"connections\": [\n      \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n    ],\n    \"created_at\": 0,\n    \"creator_id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"name\": \"string\",\n    \"servers\": [\n      \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n    ]\n  }\n}",
				ResponseBodyInfo: "The response returns the room data, including identifiers, name, connection information, creation time, and servers.",
			},
		},
	}

	getRoomsAllDoc := types.DocStruct{
		Title: "Get All Rooms",
		Info:  "This method is used to get information about all rooms in the system. The request requires a token in the header for authorization. The response returns a list of all available rooms with their data.",
		DockLists: []types.DocStructList{
			{
				Title:            "Get All Rooms",
				Description:      "Request to get information about all rooms in the system. The request must be authorized using a token.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsAll",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "The request body is empty as the data is retrieved without the need for additional parameters. You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"rooms\": [\n      {\n        \"connections\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ],\n        \"createdAt\": 0,\n        \"creatorID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"existsServers\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ],\n        \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"isExists\": true,\n        \"name\": \"string\",\n        \"roomCode\": \"string\",\n        \"servers\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ]\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "The response returns an array of objects, each representing a room. Room fields include identifiers, name, room code (roomCode), connection information, servers, and existence status.",
			},
		},
	}

	joinRoomDoc := types.DocStruct{
		Title: "Join Room by ID",
		Info:  "This method is used to join a room based on its ID. The request requires a token in the header for authorization. The response returns a new token containing room information.",
		DockLists: []types.DocStructList{
			{
				Title:            "Join Room",
				Description:      "Request to join a room by its ID. The request must be authorized using a token.",
				RequestPath:      "/api/v1/devToolsConnections/joinRoomByID",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body contains the room ID (roomID). You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "The response returns a new token that contains information about the room and the connected client.",
			},
		},
	}

	joinRoomByRoomCodeDoc := types.DocStruct{
		Title: "Join Room by Code",
		Info:  "This method is used to join a room based on its code. The request requires a token in the header for authorization. The response returns a new token containing room information.",
		DockLists: []types.DocStructList{
			{
				Title:            "Join Room by Code",
				Description:      "Request to join a room using its unique code. The request must be authorized with a token.",
				RequestPath:      "/api/v1/devToolsConnections/joinRoomByRoomCode",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomCode\": \"string\"\n  }\n}",
				RequestBodyInfo:  "The request body contains the unique room code (roomCode). You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "The response returns a new token containing information about the room and the connected client.",
			},
		},
	}

	leaveRoomDoc := types.DocStruct{
		Title: "Leave Room",
		Info:  "This method is used to exit the specified room. The request requires a token in the header for authorization.",
		DockLists: []types.DocStructList{
			{
				Title:            "Leave Room",
				Description:      "Request to exit the room using its ID. The request must be authorized with a token.",
				RequestPath:      "/api/v1/devToolsConnections/leaveRoom",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body contains the room ID (roomID) to leave. You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "The response returns an empty object confirming successful room exit.",
			},
		},
	}

	removeRoomDoc := types.DocStruct{
		Title: "Remove Room by ID",
		Info:  "This method is used to remove a room by its ID. Removal can only be performed by the room creator. The request requires a token in the header for authorization.",
		DockLists: []types.DocStructList{
			{
				Title:            "Remove Room",
				Description:      "Request to remove a room by its ID. Removal is only available for the room creator and requires authorization with a token.",
				RequestPath:      "/api/v1/devToolsConnections/removeRoomByID",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body contains the room ID (roomID) to be removed. You must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "The response does not return data, confirming successful removal operation.",
			},
		},
	}

	getRoomsConnectionUrlsDoc := types.DocStruct{
		Title: "Get Connection URLs for Room Servers",
		Info:  "This method is used to get the connection addresses for users to connect to servers after creating a room. The server distributes users across servers and returns the connection addresses. The request does not require additional parameters but requires a token in the header for authorization.",
		DockLists: []types.DocStructList{
			{
				Title:            "Get Connection URL",
				Description:      "Request to get the connection addresses for the servers to which users of the created room will be distributed. This method is similar to load balancing across servers.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsConnectionUrls",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "No parameters need to be passed in the request body, but you must provide the token in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"connectionsServer\": [\n      {\n        \"address\": \"string\",\n        \"connectionPort\": \"string\",\n        \"fullURL\": \"string\",\n        \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"isExists\": true,\n        \"path\": \"string\",\n        \"serverType\": \"string\"\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "The response returns an array of `connectionsServer` objects. Each object contains:\n- `address`: connection address\n- `connectionPort`: connection port\n- `fullURL`: full connection URL\n- `id`: unique server identifier\n- `isExists`: server availability status\n- `path`: connection path\n- `serverType`: server type (udp, tcp, websocket).",
			},
		},
	}

	return []types.DocStruct{creteRoomDoc, getMyRoomDoc, getRoomsAllDoc, joinRoomDoc, joinRoomByRoomCodeDoc, removeRoomDoc, leaveRoomDoc, getRoomsConnectionUrlsDoc}
}
