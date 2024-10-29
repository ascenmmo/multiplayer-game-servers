package ruconnections

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetConnectionsStruct() []types.DocStruct {
	creteRoomDoc := types.DocStruct{
		Title: "Создание комнаты",
		Info:  "Этот метод используется для создания новой комнаты в игре. Запрос требует наличие токена в заголовке для авторизации. В ответе возвращается новый токен, который может быть использован для последующих запросов.",
		DockLists: []types.DocStructList{
			{
				Title:            "Создать комнату",
				Description:      "Запрос на создание новой комнаты для игры. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/createRoom",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"serverType\": \"string\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передаются идентификатор игры (gameID) и тип сервера (serverType). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается новый токен, который может быть использован для последующих запросов.",
			},
		},
	}

	getRoomsDoc := types.DocStruct{
		Title: "Получение всех комнат",
		Info:  "Этот метод используется для получения информации о всех комнатах для указанной игры. Запрос требует наличие токена в заголовке для авторизации. В ответе возвращается список комнат с их данными.",
		DockLists: []types.DocStructList{
			{
				Title:            "Получить все комнаты",
				Description:      "Запрос на получение информации о всех комнатах для заданной игры. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsAll",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается идентификатор игры (gameID). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"rooms\": [\n      {\n        \"connections\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ],\n        \"created_at\": 0,\n        \"creator_id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"name\": \"string\",\n        \"servers\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ]\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается массив комнат с их данными, включая идентификаторы, имя, информацию о соединениях, времени создания и сервере.",
			},
		},
	}

	joinRoomDoc := types.DocStruct{
		Title: "Присоединение к комнате по ID",
		Info:  "Этот метод используется для присоединения к комнате на основе ее идентификатора. Запрос требует наличие токена в заголовке для авторизации. В ответе возвращается новый токен, содержащий информацию о комнате.",
		DockLists: []types.DocStructList{
			{
				Title:            "Присоединиться к комнате",
				Description:      "Запрос на присоединение к комнате по ее идентификатору. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/joinRoomByID",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передаются идентификатор игры (gameID) и идентификатор комнаты (roomID). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается новый токен, который содержит информацию о комнате и подключенном клиенте.",
			},
		},
	}

	removeRoomDoc := types.DocStruct{
		Title: "Удаление комнаты по ID",
		Info:  "Этот метод используется для удаления комнаты по ее идентификатору. Удаление может быть выполнено только создателем комнаты. Запрос требует наличие токена в заголовке для авторизации.",
		DockLists: []types.DocStructList{
			{
				Title:            "Удалить комнату",
				Description:      "Запрос на удаление комнаты по ее идентификатору. Удаление доступно только для создателя комнаты и требует авторизации с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/removeRoomByID",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передаются идентификатор игры (gameID) и идентификатор комнаты (roomID). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "В ответе не возвращаются данные, что подтверждает успешное выполнение операции удаления.",
			},
		},
	}

	getRoomsConnectionUrlsDoc := types.DocStruct{
		Title: "Получение URL для подключения к серверам комнат",
		Info:  "Этот метод используется для получения адресов подключения пользователей к серверам после создания комнаты. Сервер распределяет пользователей по серверам и возвращает адреса подключения. Запрос не требует дополнительных параметров и требует наличие токена в заголовке для авторизации.",
		DockLists: []types.DocStructList{
			{
				Title:            "Получить URL для подключения",
				Description:      "Запрос на получение адресов подключения к серверам, куда будут распределены пользователи созданной комнаты. Этот метод аналогичен балансировке нагрузки между серверами.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsConnectionUrls",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "В теле запроса не требуется передавать параметры, но необходимо указать токен в заголовке для авторизации.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"connectionsServer\": [\n      {\n        \"Address\": \"string\",\n        \"ServerType\": \"string\"\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается массив объектов `connectionsServer`, каждый из которых содержит адрес подключения (`Address`) и тип сервера (`ServerType`). Возможные значения для `ServerType`: udp, tcp, websocket.",
			},
		},
	}

	return []types.DocStruct{creteRoomDoc, getRoomsDoc, joinRoomDoc, removeRoomDoc, getRoomsConnectionUrlsDoc}
}
