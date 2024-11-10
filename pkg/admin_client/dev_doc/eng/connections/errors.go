package engconnections

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func ConnectionErrors() (errors []types.DocErrorList) {
	token := types.DocErrorList{
		Name:        "Token Expired",
		Description: "The token is no longer active. Please refresh the token using the /api/v1/devToolsClient/refreshToken endpoint or refer to the 'Clients' section for instructions.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"token is expired by 215h25m2s\"\n    }\n}",
	}

	gameNotFound := types.DocErrorList{
		Name:        "Invalid Game ID",
		Description: "The game with the specified ID was not found. Please ensure the gameID is correct and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error game not found\"\n    }\n}",
	}

	serverOff := types.DocErrorList{
		Name:        "No Available Servers for Connection",
		Description: "There are no active servers available for creating a room. Please activate the necessary servers in the admin panel and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error creating room all servers off\"\n    }\n}",
	}

	roomNotFound := types.DocErrorList{
		Name:        "Room Not Found",
		Description: "The requested room does not exist. It may have been deleted or never created. Please check the room ID and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room not found\"\n    }\n}",
	}

	accessDenied := types.DocErrorList{
		Name:        "Access to Room Denied",
		Description: "You do not have access to this room. Please ensure you are a participant of the room to gain access.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"access denied\"\n    }\n}",
	}

	errors = append(errors, token, gameNotFound, serverOff, roomNotFound, accessDenied)

	return errors
}
