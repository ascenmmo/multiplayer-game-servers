package ruconnections

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func ConnectionErrors() (errors []types.DocErrorList) {
	token := types.DocErrorList{
		Name:        "Срок действия токена истек",
		Description: "Токен больше не активен. Пожалуйста, обновите токен через ручку /api/v1/devToolsClient/refreshToken или обратитесь к разделу 'Клиенты' для инструкций.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"token is expired by 215h25m2s\"\n    }\n}",
	}

	gameNotFound := types.DocErrorList{
		Name:        "Неверный идентификатор игры",
		Description: "Игра с указанным ID не найдена. Убедитесь, что вы ввели правильный gameID, и повторите попытку.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error game not found\"\n    }\n}",
	}

	serverOff := types.DocErrorList{
		Name:        "Нет доступных серверов для подключения",
		Description: "Для создания комнаты нет активных серверов. Пожалуйста, активируйте нужные серверы в админке и повторите попытку.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"error creating room all servers off\"\n    }\n}",
	}

	roomNotFound := types.DocErrorList{
		Name:        "Комната не найдена",
		Description: "Запрашиваемая комната не существует. Возможно, она была удалена или не была создана. Проверьте корректность ID комнаты и повторите запрос.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room not found\"\n    }\n}",
	}

	accessDenied := types.DocErrorList{
		Name:        "Доступ к комнате запрещен",
		Description: "Вы не имеете доступа к этой комнате. Убедитесь, что вы являетесь участником комнаты, чтобы получить к ней доступ.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"access denied\"\n    }\n}",
	}

	errors = append(errors, token, gameNotFound, serverOff, roomNotFound, accessDenied)

	return errors
}
