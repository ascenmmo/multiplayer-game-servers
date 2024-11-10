package engconnections

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetConnectionsStruct() []types.DocStruct {
	createRoomDoc := types.DocStruct{
		Title: "Create Room",
		Info:  "This method is used to create a new room in the game. The request requires a token in the header for authorization. The response returns a new token, which can be used for subsequent requests.",
		DockLists: []types.DocStructList{
			{
				Title:            "Create Room",
				Description:      "Request to create a new game room. The request must be authorized with a token.",
				RequestPath:      "/api/v1/devToolsConnections/createRoom",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"serverType\": \"string\"\n  }\n}",
				RequestBodyInfo:  "The request body includes the game ID (gameID) and server type (serverType). A token is required in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "The response includes a new token that can be used for future requests.",
			},
		},
	}

	getRoomsDoc := types.DocStruct{
		Title: "Get All Rooms",
		Info:  "This method is used to retrieve information about all rooms for the specified game. The request requires a token in the header for authorization. The response returns a list of rooms with their details.",
		DockLists: []types.DocStructList{
			{
				Title:            "Get All Rooms",
				Description:      "Request to get information about all rooms for a given game. The request must be authorized with a token.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsAll",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body includes the game ID (gameID). A token is required in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"rooms\": [\n      {\n        \"connections\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ],\n        \"created_at\": 0,\n        \"creator_id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"name\": \"string\",\n        \"servers\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ]\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "The response returns an array of rooms with details, including IDs, names, connections, creation time, and server information.",
			},
		},
	}

	joinRoomDoc := types.DocStruct{
		Title: "Join Room by ID",
		Info:  "This method is used to join a room based on its ID. The request requires a token in the header for authorization. The response returns a new token containing room information.",
		DockLists: []types.DocStructList{
			{
				Title:            "Join Room",
				Description:      "Request to join a room by its ID. The request must be authorized with a token.",
				RequestPath:      "/api/v1/devToolsConnections/joinRoomByID",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body includes the game ID (gameID) and room ID (roomID). A token is required in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "The response returns a new token containing information about the room and connected client.",
			},
		},
	}

	removeRoomDoc := types.DocStruct{
		Title: "Remove Room by ID",
		Info:  "This method is used to delete a room by its ID. Only the room creator can delete the room. The request requires a token in the header for authorization.",
		DockLists: []types.DocStructList{
			{
				Title:            "Delete Room",
				Description:      "Request to delete a room by its ID. Only the room creator can delete the room and must be authorized with a token.",
				RequestPath:      "/api/v1/devToolsConnections/removeRoomByID",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body includes the game ID (gameID) and room ID (roomID). A token is required in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "The response does not return data, indicating successful deletion of the room.",
			},
		},
	}

	getRoomsConnectionUrlsDoc := types.DocStruct{
		Title: "Get Room Connection URLs",
		Info:  "This method is used to get connection addresses for users to servers after creating a room. The server distributes users across servers and returns connection addresses. The request requires a token in the header for authorization.",
		DockLists: []types.DocStructList{
			{
				Title:            "Get Connection URLs",
				Description:      "Request to get connection addresses to servers where users are distributed after room creation. This method functions similarly to load balancing between servers.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsConnectionUrls",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "No parameters are needed in the request body, but a token is required in the header for authorization.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"connectionsServer\": [\n      {\n        \"Address\": \"string\",\n        \"ServerType\": \"string\"\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "The response returns an array of `connectionsServer` objects, each containing a connection address (`Address`) and server type (`ServerType`). Possible values for `ServerType`: udp, tcp, websocket.",
			},
		},
	}

	return []types.DocStruct{createRoomDoc, getRoomsDoc, joinRoomDoc, removeRoomDoc, getRoomsConnectionUrlsDoc}
}
