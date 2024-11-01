package ruclient

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func ClientErrors() (errors []types.DocErrorList) {
	regGame := types.DocErrorList{
		Name:        "Ошибка при регистрации игры",
		Description: "Игра с таким ID не найдена. Проверьте введенный идентификатор игры и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error game not found\"\n    }\n}",
	}

	regClient := types.DocErrorList{
		Name:        "Ошибка при регистрации клиента",
		Description: "Клиент с таким именем уже существует. Пожалуйста, используйте другое имя.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error creating client\"\n    }\n}",
	}

	authClient := types.DocErrorList{
		Name:        "Ошибка при авторизации клиента",
		Description: "Клиент с таким идентификатором не найден. Проверьте правильность данных для входа и повторите попытку.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error client not found\"\n    }\n}",
	}

	errors = append(errors, regGame, regClient, authClient)
	return errors
}
