package engclient

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func ClientErrors() (errors []types.DocErrorList) {
	regGame := types.DocErrorList{
		Name:        "Error when registering a game",
		Description: "Game with the specified ID not found. Please check the entered game ID and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error game not found\"\n    }\n}",
	}

	regClient := types.DocErrorList{
		Name:        "Error when registering a client",
		Description: "A client with this name already exists. Please use a different name.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error creating client\"\n    }\n}",
	}

	authClient := types.DocErrorList{
		Name:        "Error when authorizing a client",
		Description: "Client with the specified ID not found. Check the login credentials and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error client not found\"\n    }\n}",
	}

	errors = append(errors, regGame, regClient, authClient)
	return errors
}
