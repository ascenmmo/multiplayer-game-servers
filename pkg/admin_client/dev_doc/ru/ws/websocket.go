package ws

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"
)

func GetWsConnectionsStruct() []types.DocStruct {
	ws := types.DocStruct{
		Title: "Подключение по WebSocket",
		Info: "Этот WebSocket-сервис разработан специально для многопользовательских онлайн игр, обеспечивая передачу данных в режиме реального времени между клиентами. " +
			"С его помощью можно быстро и эффективно обмениваться игровыми данными, что особенно важно в условиях, где скорость имеет приоритет над надёжностью доставки сообщений. " +
			"WebSocket поддерживает двустороннюю связь между клиентом и сервером, что делает его отличным выбором для игровых клиентов. Однако следует учитывать, что возможны потери пакетов или их получение в неправильном порядке.",
		DockLists: []types.DocStructList{
			{
				Title: "Отправка и приём сообщений",
				Description: "Для подключения к игровому WebSocket-сервису вам нужно знать URL и использовать токен для аутентификации. " +
					"Максимальное количество запросов в секунду — 200." +
					"Пример команды для подключения и отправки игрового сообщения через WebSocket:",
				RequestPath:   "ws://127.0.0.1:4240/api/ws/connect",
				RequestHeader: "Token: ваш токен",
				Method:        "ws",
				RequestBody: `{
	"data": "данные для отправки"
}`,
				RequestBodyInfo: "Тело запроса должно содержать игровую информацию, отправляемую другим игрокам. Токен для аутентификации передается в заголовке.",
				ResponseBody: `{
	"data": "данные от других клиентов"
}`,
				ResponseBodyInfo: "Ответ содержит данные, полученные от других игровых клиентов, подключённых к той же игровой комнате.",
			},
		},
	}

	return []types.DocStruct{ws}
}
