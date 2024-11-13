package engclient

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetClientStruct() []types.DocStruct {
	auth := types.DocStruct{
		Title: "User Authentication",
		Info: "Each user must be authenticated to receive a unique token. " +
			"Servers will not accept requests if the token is missing or invalid. " +
			"The token also contains embedded information about the game and the player.",
		DockLists: []types.DocStructList{
			{
				Title: "Registration",
				Description: "Register a new user in the system by providing a unique email and password. " +
					"Upon successful registration, tokens are returned.",
				RequestPath:   "https://ascenmmo.com/api/v1/devToolsClient/signUp",
				Method:        "POST",
				RequestHeader: "Content-Type: application/json",
				RequestBody:   "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n      \"password\": \"string\"\n    }\n  }\n}",
				RequestBodyInfo: "The 'email' field must be unique. If the email is reused, an error is returned. " +
					"Additional information can be passed in JSON format in the 'additional' field.",
				ResponseBody: "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"refresh\": \"string\",\n    \"token\": \"string\"\n  }\n}",
				ResponseBodyInfo: "Upon successful registration, two fields are returned: 'token' and 'refresh'. " +
					"These tokens should be saved for subsequent use in request headers.",
			},
			{
				Title:            "Authentication",
				Description:      "Authenticate a user using email and password. Tokens are returned upon successful authentication.",
				RequestPath:      "https://ascenmmo.com/api/v1/devToolsClient/signIn",
				Method:           "POST",
				RequestHeader:    "Content-Type: application/json",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n      \"password\": \"string\"\n    }\n  }\n}",
				RequestBodyInfo:  "Provide email and password for authentication. Tokens are returned upon success.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"refresh\": \"string\",\n    \"token\": \"string\"\n  }\n}",
				ResponseBodyInfo: "Returns 'token' and 'refresh' which must be used for further requests.",
			},
		},
	}

	refreshTokenDoc := types.DocStruct{
		Title: "Token Refresh",
		Info: "This method is used to refresh access and refresh tokens. " +
			"Current tokens must be passed in the request headers. No additional parameters are required in the request body.",
		DockLists: []types.DocStructList{
			{
				Title:           "Refresh Tokens",
				Description:     "Request to refresh access and refresh tokens. Requires token and refresh token in headers.",
				RequestPath:     "/api/v1/devToolsClient/refreshToken",
				Method:          "POST",
				RequestHeader:   "Token: your token\nRefreshToken: your refresh token",
				RequestBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo: "No additional parameters in the request body, just pass an empty object in 'params'.",
				ResponseBody:    "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newRefresh\": \"string\",\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "The response includes new tokens: 'newRefresh' and 'newToken'. " +
					"These tokens should be saved and used for future requests.",
			},
		},
	}

	clientDoc := types.DocStruct{
		Title: "Retrieve Client Data",
		Info: "This method is used to retrieve client information based on the game ID (gameID). " +
			"A token must be included in the header for authentication. The response returns client data such as email, nickname, gameID, and additional information.",
		DockLists: []types.DocStructList{
			{
				Title:            "Get Client Data",
				Description:      "Request to retrieve client information by gameID. Must be authorized using a token.",
				RequestPath:      "/api/v1/devToolsClient/getClient",
				Method:           "POST",
				RequestHeader:    "Token: your token",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "The request body contains the game ID (gameID). A token must be included in the header to authorize the request.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n    }\n  }\n}",
				ResponseBodyInfo: "The response includes a client object containing email, gameID, nickname, and additional information in the 'additional' field.",
			},
		},
	}

	updateClientDoc := types.DocStruct{
		Title: "Update Client Data",
		Info: "This method is used to update client information. " +
			"A token must be passed in the request header and a client object in the request body. " +
			"The client object can include a new password if needed.",
		DockLists: []types.DocStructList{
			{
				Title:         "Update Client Data",
				Description:   "Request to update client data. Requires a token in the header and a client object in the body.",
				RequestPath:   "/api/v1/devToolsClient/updateClient",
				Method:        "POST",
				RequestHeader: "Token: your token",
				RequestBody:   "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"nickname\": \"string\",\n      \"password\": \"string\",\n      \"newPassword\": \"string\"\n    }\n  }\n}",
				RequestBodyInfo: "The request body includes a client object with the following fields: \n" +
					"- `additional`: additional information (can be an empty object)\n" +
					"- `email`: new client email\n" +
					"- `nickname`: new client nickname\n" +
					"- `password`: current client password (for verification)\n" +
					"- `newPassword`: new client password (if required).",
				ResponseBody: "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "Upon successful update, the response includes: \n" +
					"- `id`: the request ID, matching the sent ID.\n" +
					"- `jsonrpc`: the JSON-RPC version, matching the sent version.\n" +
					"- `result`: an empty object confirming the successful request. " +
					"This object indicates that the client data has been successfully updated.",
			},
		},
	}

	getGameSavesDoc := types.DocStruct{
		Title: "Get Game Saves",
		Info: "This method is used to retrieve the game saves. " +
			"You need to pass a token in the request header and an empty request body in JSON-RPC format.",
		DockLists: []types.DocStructList{
			{
				Title:         "Get Game Saves",
				Description:   "Request to get game saves. Token in the header and an empty body in the request are required.",
				RequestPath:   "/api/v1/devToolsClient/getGameSaves",
				Method:        "POST",
				RequestHeader: "Token: your token",
				RequestBody:   "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo: "The request body contains an empty `params` object in JSON-RPC format:\n" +
					"- `id`: request identifier\n" +
					"- `jsonrpc`: version of JSON-RPC\n" +
					"- `params`: empty object",
				ResponseBody: "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"gameSaves\": {\n      \"saves\": {}\n    }\n  }\n}",
				ResponseBodyInfo: "If the request is successful, it returns an object with the following fields:\n" +
					"- `id`: request identifier, matching the sent one\n" +
					"- `jsonrpc`: version of JSON-RPC\n" +
					"- `result`: object containing the game saves\n" +
					"  - `gameSaves`: object with game save data",
			},
		},
	}

	setGameSavesDoc := types.DocStruct{
		Title: "Save Game Data",
		Info: "This method is used to save game data. " +
			"You need to pass a token in the request header and an object with game saves in the request body.",
		DockLists: []types.DocStructList{
			{
				Title:         "Save Game Data",
				Description:   "Request to save game data. Token in the header and an object with game saves in the request body are required.",
				RequestPath:   "/api/v1/devToolsClient/setGameSaves",
				Method:        "POST",
				RequestHeader: "Token: your token",
				RequestBody:   "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameSaves\": {\n      \"saves\": {\n        \"FirstSave\": {\"location\": \"loc1\", \"health\": 100}\n      }\n    }\n  }\n}",
				RequestBodyInfo: "The request body contains the `gameSaves` object, which includes:\n" +
					"- `saves`: an object with game save data. Example:\n" +
					"  - `FirstSave`: contains a save with fields:\n" +
					"    - `location`: the location of the save\n" +
					"    - `health`: the health state",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "If the request is successful, it returns an empty `result` object indicating successful save.",
			},
		},
	}

	deleteGameSavesDoc := types.DocStruct{
		Title: "Delete Game Data",
		Info: "This method is used to delete game save data. " +
			"You need to pass a token in the request header and an empty request body in JSON-RPC format.",
		DockLists: []types.DocStructList{
			{
				Title:         "Delete Game Saves",
				Description:   "Request to delete game saves. Token in the header and an empty request body are required.",
				RequestPath:   "/api/v1/devToolsClient/deleteGameSaves",
				Method:        "POST",
				RequestHeader: "Token: your token",
				RequestBody:   "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo: "The request body contains an empty `params` object in JSON-RPC format:\n" +
					"- `id`: request identifier\n" +
					"- `jsonrpc`: version of JSON-RPC\n" +
					"- `params`: empty object",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "If the request is successful, it returns an empty `result` object confirming that the saves have been deleted.",
			},
		},
	}

	return []types.DocStruct{auth, refreshTokenDoc, clientDoc, updateClientDoc, getGameSavesDoc, setGameSavesDoc, deleteGameSavesDoc}

}
