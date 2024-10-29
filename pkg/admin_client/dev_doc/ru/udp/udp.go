package udp

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func GetUPDConnectionsStruct() []types.DocStruct {
	udp := types.DocStruct{
		Title: "Подключение по UDP",
		Info: "Этот сервис предназначен для передачи небольших данных в режиме реального времени. " +
			"Если ваши клиенты хотят получать данные максимально быстро, а скорость важнее, чем гарантированная доставка, то этот сервис подходит идеально. " +
			"UDP — это очень быстрое решение, однако, есть недостатки: некоторые пакеты могут быть потеряны или доставлены не по порядку.",
		DockLists: []types.DocStructList{
			{
				Title: "Отправка и приём сообщений",
				Description: "Для подключения к сервису необходимо знать IP-адрес и порт. " +
					"Максимальное количество сообщений в секунду — 200. " +
					"Пример команды подключения и отправки сообщения через UDP:",
				RequestPath: "udp://127.0.0.1:4500",
				Method:      "udp",
				RequestBody: `{
	"token": "ваш токен", 
	"data": "данные для отправки"
}`,
				RequestBodyInfo: "Тело запроса содержит токен для аутентификации и данные, которые необходимо передать другим клиентам.",
				ResponseBody: `{
	"data": "данные от других клиентов"
}`,
				ResponseBodyInfo: "Ответ содержит данные, полученные от других клиентов, подключённых к той же комнате.",
				ErrorBody: `{
	"error": "информация об ошибке"
}`,
				DocErrorList: []types.DocErrorList{
					{
						Name:        "room not found",
						Description: "Ошибка возникает, если комната, к которой пытается подключиться клиент, не существует.",
						Body:        `{"error": "room not found"}`,
					},
					{
						Name:        "room bad value",
						Description: "Ошибка возникает, если переданы неверные параметры для подключения к комнате.",
						Body:        `{"error": "room bad value"}`,
					},
					{
						Name:        "too many connections",
						Description: "Ошибка возникает, если превышено максимальное количество подключений или сообщений в секунду.",
						Body:        `{"error": "too many requests"}`,
					},
				},
			},
		},
	}

	return []types.DocStruct{udp}
}
