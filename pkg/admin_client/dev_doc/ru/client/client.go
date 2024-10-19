package ruclient

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetClientStruct() []types.DocStruct {
	auth := types.DocStruct{
		Title: "Авторизация вашего пользователя",
		Info: "Каждый пользователь должен быть авторизирован для получения уникального токена." +
			"Серверы не пропустят запрос, если токена нет или он неверный." +
			"Также в токен вшита информация об игре и игроке.",
		DockLists: []types.DocStructList{
			{
				Title: "Регистрация",
				Description: "Регистрация нового пользователя в системе с указанием уникального email и пароля. " +
					"После успешной регистрации возвращаются токены.",
				RequestPath:   "https://asenmmo.com/api/v1/devToolsClient/signUp",
				Method:        "POST",
				RequestHeader: "Content-Type: application/json",
				ResponseBody:  "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n      \"password\": \"string\"\n    }\n  }\n}",
				RequestBodyInfo: "Поле 'email' должно быть уникальным. Если email используется повторно, возвращается ошибка. " +
					"Дополнительную информацию можно передавать в формате JSON в поле 'additional'.",
				RequestBody: "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"refresh\": \"string\",\n    \"token\": \"string\"\n  }\n}",
				ResponseBodyInfo: "При успешной регистрации возвращаются два поля: 'token' и 'refresh'. " +
					"Эти токены следует сохранить для последующего использования в заголовках запросов.",
			},
			{
				Title:            "Авторизация",
				Description:      "Авторизация пользователя с использованием email и пароля. После успешной авторизации возвращаются токены.",
				RequestPath:      "https://asenmmo.com/api/v1/devToolsClient/signIn",
				Method:           "POST",
				RequestHeader:    "Content-Type: application/json",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n      \"password\": \"string\"\n    }\n  }\n}",
				RequestBodyInfo:  "Передайте email и пароль для авторизации. В случае успеха вернутся токены.",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"refresh\": \"string\",\n    \"token\": \"string\"\n  }\n}",
				ResponseBodyInfo: "Возвращаются 'token' и 'refresh', которые необходимо использовать для дальнейших запросов.",
				ErrorBody:        "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\"}",
				DocErrorList: []types.DocErrorList{{
					Name:        "error",
					Description: "если пользователь не найден",
					Body:        "'error'': {'user'}",
				},
				},
			},
		},
	}

	refreshTokenDoc := types.DocStruct{
		Title: "Обновление токенов",
		Info: "Данный метод используется для обновления токенов доступа и обновления. " +
			"Необходимо передать текущие токены в заголовках запроса. В теле запроса нет дополнительных параметров.",
		DockLists: []types.DocStructList{
			{
				Title:           "Обновить токены",
				Description:     "Запрос на обновление токенов доступа и обновления. Требуется токен и refresh токен в заголовках.",
				RequestPath:     "/api/v1/devToolsClient/refreshToken",
				Method:          "POST",
				RequestHeader:   "Token: ваш токен\nRefreshToken: ваш refresh токен",
				RequestBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo: "В теле запроса нет дополнительных параметров, просто передайте пустой объект в 'params'.",
				ResponseBody:    "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newRefresh\": \"string\",\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "В ответе возвращаются новые токены: 'newRefresh' и 'newToken'. " +
					"Эти токены следует сохранить и использовать для дальнейших запросов.",
			},
		},
	}

	clientDoc := types.DocStruct{
		Title: "Получение данных клиента",
		Info: "Этот метод используется для получения информации о клиенте на основе идентификатора игры (gameID). " +
			"Запрос требует наличие токена в заголовке для авторизации. В ответе возвращаются данные клиента, такие как email, nickname, gameID, а также дополнительная информация.",
		DockLists: []types.DocStructList{
			{
				Title:            "Получить данные клиента",
				Description:      "Запрос на получение информации о клиенте по его gameID. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsClient/getClient",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается идентификатор игры (gameID). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n    }\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается объект клиента, который содержит email, gameID, nickname и дополнительную информацию в поле 'additional'.",
			},
		},
	}

	updateClientDoc := types.DocStruct{
		Title: "Обновление данных клиента",
		Info: "Данный метод используется для обновления информации о клиенте. " +
			"Необходимо передать токен в заголовке запроса и объект клиента в теле запроса. " +
			"Объект клиента может включать новый пароль, если он требуется.",
		DockLists: []types.DocStructList{
			{
				Title:         "Обновить данные клиента",
				Description:   "Запрос на обновление данных клиента. Требуется токен в заголовке и объект клиента в теле.",
				RequestPath:   "/api/v1/devToolsClient/updateClient",
				Method:        "POST",
				RequestHeader: "Token: ваш токен",
				RequestBody:   "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"nickname\": \"string\",\n      \"password\": \"string\",\n      \"newPassword\": \"string\"\n    }\n  }\n}",
				RequestBodyInfo: "В теле запроса передается объект клиента, который включает поля: \n" +
					"- `additional`: дополнительная информация (может быть пустым объектом)\n" +
					"- `email`: новый email клиента\n" +
					"- `nickname`: новый nickname клиента\n" +
					"- `password`: текущий пароль клиента (для проверки)\n" +
					"- `newPassword`: новый пароль клиента (если требуется изменить).",
				ResponseBody: "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "При успешном обновлении возвращается объект, содержащий поля:\n" +
					"- `id`: идентификатор запроса, который совпадает с тем, что был отправлен.\n" +
					"- `jsonrpc`: версия JSON-RPC, которая также совпадает с той, что была отправлена.\n" +
					"- `result`: пустой объект, подтверждающий успешное выполнение запроса. " +
					"Этот объект указывает на то, что данные клиента были успешно обновлены.",
			},
		},
	}
	return []types.DocStruct{auth, refreshTokenDoc, clientDoc, updateClientDoc}
}
