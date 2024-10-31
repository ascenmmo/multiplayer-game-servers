package tcp

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func TcpErrors() (errorsList []types.DocErrorList) {
	tokenError := types.DocErrorList{
		Name:        "Истек срок действия токена",
		Description: "Ваш токен больше не действителен. Пожалуйста, обновите его через /api/v1/devToolsClient/refreshToken или обратитесь к разделу 'Клиенты' для инструкций.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"token is expired by 215h25m2s\"\n    }\n}",
	}

	userNotFoundError := types.DocErrorList{
		Name:        "Пользователь не найден",
		Description: "Система не обнаружила указанного пользователя. Убедитесь, что вы ввели правильные данные и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"user not found\"\n    }\n}",
	}

	roomNotFoundError := types.DocErrorList{
		Name:        "Комната не найдена",
		Description: "Система не обнаружила запрашиваемую комнату. Проверьте идентификатор комнаты и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room not found\"\n    }\n}",
	}

	roomIsExistsError := types.DocErrorList{
		Name:        "Комната уже существует",
		Description: "Попытка создать комнату, которая уже есть в системе. Проверьте настройки и повторите запрос.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room is exists\"\n    }\n}",
	}

	roomBadValueError := types.DocErrorList{
		Name:        "Некорректное значение для комнаты",
		Description: "Введенное значение для комнаты недопустимо. Проверьте данные и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"room bad value\"\n    }\n}",
	}

	tooManyRequestsError := types.DocErrorList{
		Name:        "Превышено количество запросов",
		Description: "Отправлено слишком много запросов. Подождите некоторое время и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"too many connection\", \"message\": \"too many requests\"\n   }\n}",
	}

	notifyServerNotFoundError := types.DocErrorList{
		Name:        "Сервер уведомлений не найден",
		Description: "Не удалось найти указанный сервер уведомлений. Проверьте настройки сервера и повторите попытку.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"err notify server not found\"\n    }\n}",
	}

	notifyServerNotValidError := types.DocErrorList{
		Name:        "Недействительный сервер уведомлений",
		Description: "Указанный сервер уведомлений недействителен. Проверьте его конфигурацию и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"err notify server not valid\"\n    }\n}",
	}

	gameConfigMarshalUserDataError := types.DocErrorList{
		Name:        "Ошибка обработки данных пользователя",
		Description: "Произошла ошибка при обработке данных пользователя в конфигурации игры. Проверьте данные и попробуйте снова.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"err game config marshal user data\"\n    }\n}",
	}

	gameResultsNotFoundError := types.DocErrorList{
		Name:        "Результаты игры не найдены",
		Description: "Система не обнаружила результаты для указанной игры. Убедитесь, что игра завершена, и повторите запрос.",
		Body:        "{\n    \"id\": 1,\n    \"jsonrpc\": \"2.0\",\n    \"error\": {\n        \"message\": \"game results not found\"\n    }\n}",
	}

	errorsList = append(errorsList, tokenError, userNotFoundError, roomNotFoundError, roomIsExistsError, roomBadValueError, tooManyRequestsError, notifyServerNotFoundError, notifyServerNotValidError, gameConfigMarshalUserDataError, gameResultsNotFoundError)

	return errorsList
}
