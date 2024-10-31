package tcp

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetTCPConnectionsStruct() []types.DocStruct {
	tcp := types.DocStruct{
		Title: "Подключение по TCP",
		Info: "Этот сервис предназначен для взаимодействия между игровыми клиентами и сервером через TCP соединение. " +
			"С его помощью клиенты могут отправлять сообщения, получать новые сообщения и удалять пользователей. " +
			"Сервис поддерживает до **10 запросов в секунду** и идеально подходит для игровых приложений, " +
			"где требуется высокая скорость взаимодействия и поддержка лонгпулинга для получения новых сообщений.",
		DockLists: []types.DocStructList{
			{
				Title:         "Отправка сообщений",
				Description:   "Для отправки сообщения от клиента на сервер используйте следующую ручку. Необходимо передать токен в заголовке запроса.",
				RequestPath:   "/api/v1/rest/gameConnections/setSendMessage",
				RequestHeader: "Token: ваш токен",
				Method:        "POST",
				RequestBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "params": {
    "message": {
      "data": "ваши данные",
      "server": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "token": "ваш токен"
    }
  }
}`,
				RequestBodyInfo: "Тело запроса должно содержать данные сообщения, идентификатор сервера и токен для аутентификации.",
				ResponseBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {}
}`,
				ResponseBodyInfo: "Ответ указывает на успешное выполнение операции.",
			},
			{
				Title:         "Получение новых сообщений",
				Description:   "Для получения новых сообщений с сервера используйте следующую ручку. Можно использовать лонгпулинг.",
				RequestPath:   "/api/v1/rest/gameConnections/getMessage",
				RequestHeader: "Token: ваш токен",
				Method:        "POST",
				RequestBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "params": {}
}`,
				RequestBodyInfo: "Тело запроса не требует параметров, необходимо только указать токен в заголовке.",
				ResponseBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {
    "messages": {
      "dataArray": [
        "новое сообщение"
      ]
    }
  }
}`,
				ResponseBodyInfo: "Ответ содержит массив новых сообщений от сервера.",
			},
			{
				Title:         "Удаление пользователя",
				Description:   "Для удаления пользователя из игрового сервера используйте следующую ручку.",
				RequestPath:   "/api/v1/rest/gameConnections/removeUser",
				RequestHeader: "Token: ваш токен",
				Method:        "POST",
				RequestBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "params": {
    "userID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
  }
}`,
				RequestBodyInfo: "Тело запроса должно содержать уникальный идентификатор пользователя, которого необходимо удалить.",
				ResponseBody: `{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {}
}`,
				ResponseBodyInfo: "Ответ указывает на успешное выполнение операции.",
			},
		},
	}

	return []types.DocStruct{tcp}
}
