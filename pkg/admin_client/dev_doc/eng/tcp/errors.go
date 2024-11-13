package engtcp

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func TcpErrors() (errorsList []types.DocErrorList) {
	tokenError := types.DocErrorList{
		Name:        "Token Expired",
		Description: "Your token is no longer valid. Please refresh it using /api/v1/devToolsClient/refreshToken or refer to the 'Clients' section for instructions.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"token is expired by 215h25m2s\"\n    }\n}",
	}

	userNotFoundError := types.DocErrorList{
		Name:        "User Not Found",
		Description: "The specified user was not found. Please ensure the provided details are correct and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"user not found\"\n    }\n}",
	}

	roomNotFoundError := types.DocErrorList{
		Name:        "Room Not Found",
		Description: "The requested room was not found. Please check the room ID and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room not found\"\n    }\n}",
	}

	roomIsExistsError := types.DocErrorList{
		Name:        "Room Already Exists",
		Description: "An attempt was made to create a room that already exists in the system. Please check your settings and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room is exists\"\n    }\n}",
	}

	roomBadValueError := types.DocErrorList{
		Name:        "Invalid Room Value",
		Description: "The provided value for the room is invalid. Please check the input and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room bad value\"\n    }\n}",
	}

	tooManyRequestsError := types.DocErrorList{
		Name:        "Too Many Requests",
		Description: "Too many requests have been sent. Please wait for a while and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"too many connection\", \"message\": \"too many requests\"\n    }\n}",
	}

	notifyServerNotFoundError := types.DocErrorList{
		Name:        "Notification Server Not Found",
		Description: "The specified notification server could not be found. Please check the server settings and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"err notify server not found\"\n    }\n}",
	}

	notifyServerNotValidError := types.DocErrorList{
		Name:        "Invalid Notification Server",
		Description: "The specified notification server is invalid. Please verify its configuration and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"err notify server not valid\"\n    }\n}",
	}

	gameConfigMarshalUserDataError := types.DocErrorList{
		Name:        "User Data Processing Error",
		Description: "An error occurred while processing user data in the game configuration. Please check the data and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"err game config marshal user data\"\n    }\n}",
	}

	gameResultsNotFoundError := types.DocErrorList{
		Name:        "Game Results Not Found",
		Description: "The system could not find the results for the specified game. Please ensure the game has concluded and try again.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"game results not found\"\n    }\n}",
	}

	errorsList = append(errorsList, tokenError, userNotFoundError, roomNotFoundError, roomIsExistsError, roomBadValueError, tooManyRequestsError, notifyServerNotFoundError, notifyServerNotValidError, gameConfigMarshalUserDataError, gameResultsNotFoundError)

	return errorsList
}
