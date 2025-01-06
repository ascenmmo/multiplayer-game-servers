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
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"name\": \"string\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается имя комнаты (name). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается новый токен, который может быть использован для последующих запросов.",
			},
		},
	}

	getMyRoomDoc := types.DocStruct{
		Title: "Получение информации о своей комнате",
		Info:  "Этот метод используется для получения информации о комнате текущего пользователя. Запрос требует наличие токена в заголовке для авторизации. В ответе возвращаются данные комнаты, включая информацию о соединениях, времени создания, серверах и другие параметры.",
		DockLists: []types.DocStructList{
			{
				Title:            "Получить свою комнату",
				Description:      "Запрос на получение информации о комнате текущего пользователя. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/getMyRoom",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "Тело запроса пустое, так как данные извлекаются для текущего авторизованного пользователя. Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"connections\": [\n      \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n    ],\n    \"created_at\": 0,\n    \"creator_id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n    \"name\": \"string\",\n    \"servers\": [\n      \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n    ]\n  }\n}",
				ResponseBodyInfo: "В ответе возвращаются данные комнаты, включая идентификаторы, имя, информацию о соединениях, времени создания и серверах.",
			},
		},
	}

	getRoomsAllDoc := types.DocStruct{
		Title: "Получение всех комнат",
		Info:  "Этот метод используется для получения информации о всех комнатах в системе. Запрос требует наличие токена в заголовке для авторизации. В ответе возвращается список всех доступных комнат с их данными.",
		DockLists: []types.DocStructList{
			{
				Title:            "Получить все комнаты",
				Description:      "Запрос на получение информации о всех комнатах в системе. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsAll",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "Тело запроса пустое, так как данные извлекаются без необходимости указания дополнительных параметров. Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"rooms\": [\n      {\n        \"connections\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ],\n        \"createdAt\": 0,\n        \"creatorID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"existsServers\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ],\n        \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"isExists\": true,\n        \"name\": \"string\",\n        \"roomCode\": \"string\",\n        \"servers\": [\n          \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n        ]\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается массив объектов, каждый из которых представляет комнату. Поля комнаты включают идентификаторы, имя, код комнаты (roomCode), информацию о соединениях, серверах и статусе существования.",
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
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается идентификатор комнаты (roomID). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается новый токен, который содержит информацию о комнате и подключенном клиенте.",
			},
		},
	}

	joinRoomByRoomCodeDoc := types.DocStruct{
		Title: "Присоединение к комнате по коду",
		Info:  "Этот метод используется для присоединения к комнате на основе ее кода. Запрос требует наличие токена в заголовке для авторизации. В ответе возвращается новый токен, содержащий информацию о комнате.",
		DockLists: []types.DocStructList{
			{
				Title:            "Присоединиться к комнате по коду",
				Description:      "Запрос на присоединение к комнате по ее уникальному коду. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/joinRoomByRoomCode",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomCode\": \"string\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается уникальный код комнаты (roomCode). Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"newToken\": \"string\"\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается новый токен, который содержит информацию о комнате и подключенном клиенте.",
			},
		},
	}

	leaveRoomDoc := types.DocStruct{
		Title: "Покинуть комнату",
		Info:  "Этот метод используется для выхода из указанной комнаты. Запрос требует наличие токена в заголовке для авторизации.",
		DockLists: []types.DocStructList{
			{
				Title:            "Покинуть комнату",
				Description:      "Запрос на выход из комнаты по ее идентификатору. Запрос должен быть авторизован с помощью токена.",
				RequestPath:      "/api/v1/devToolsConnections/leaveRoom",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается идентификатор комнаты (roomID), которую необходимо покинуть. Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "В ответе возвращается пустой объект, подтверждающий успешный выход из комнаты.",
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
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"roomID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"\n  }\n}",
				RequestBodyInfo:  "В теле запроса передается идентификатор комнаты (roomID), которую необходимо удалить. Необходимо указать токен в заголовке для авторизации запроса.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {}\n}",
				ResponseBodyInfo: "В ответе не возвращаются данные, что подтверждает успешное выполнение операции удаления.",
			},
		},
	}

	getRoomsConnectionUrlsDoc := types.DocStruct{
		Title: "Получение URL для подключения к серверам комнат",
		Info:  "Этот метод используется для получения адресов подключения пользователей к серверам после создания комнаты. Сервер распределяет пользователей по серверам и возвращает адреса подключения. Запрос не требует дополнительных параметров, но требует наличие токена в заголовке для авторизации.",
		DockLists: []types.DocStructList{
			{
				Title:            "Получить URL для подключения",
				Description:      "Запрос на получение адресов подключения к серверам, куда будут распределены пользователи созданной комнаты. Этот метод аналогичен балансировке нагрузки между серверами.",
				RequestPath:      "/api/v1/devToolsConnections/getRoomsConnectionUrls",
				Method:           "POST",
				RequestHeader:    "Token: ваш токен",
				RequestBody:      "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {}\n}",
				RequestBodyInfo:  "В теле запроса не требуется передавать параметры, но необходимо указать токен в заголовке для авторизации.",
				ResponseBody:     "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"connectionsServer\": [\n      {\n        \"address\": \"string\",\n        \"connectionPort\": \"string\",\n        \"fullURL\": \"string\",\n        \"id\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n        \"isExists\": true,\n        \"path\": \"string\",\n        \"serverType\": \"string\"\n      }\n    ]\n  }\n}",
				ResponseBodyInfo: "В ответе возвращается массив объектов `connectionsServer`. Каждый объект содержит:\n- `address`: адрес подключения\n- `connectionPort`: порт подключения\n- `fullURL`: полный URL подключения\n- `id`: уникальный идентификатор сервера\n- `isExists`: статус доступности сервера\n- `path`: путь подключения\n- `serverType`: тип сервера (udp, tcp, websocket).",
			},
		},
	}

	return []types.DocStruct{creteRoomDoc, getMyRoomDoc, getRoomsAllDoc, joinRoomDoc, joinRoomByRoomCodeDoc, removeRoomDoc, leaveRoomDoc, getRoomsConnectionUrlsDoc}
}
