package ruclient

const (
	authTitle = "Авторизация вашего пользователя"
	authInfo  = "Каждый пользователь должен быть авторизирован для получения уникального токена." +
		"Серверы не пропустят запрос, если токена нет или он неверный." +
		"Также в токен вшита информация об игре и игроке."

	registrationTitle           = "Регистрация"
	registrationDescription     = "Регистрация нового пользователя в системе с указанием уникального email и пароля. После успешной регистрации возвращаются токены."
	registrationRequestPath     = "https://asenmmo.com/api/v1/devToolsClient/signUp"
	registrationRequestHeader   = "Content-Type: application/json"
	registrationResponseBody    = "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n      \"password\": \"string\"\n    }\n  }\n}"
	registrationRequestBodyInfo = "Поле 'email' должно быть уникальным. Если email используется повторно, возвращается ошибка." +
		"Дополнительную информацию можно передавать в формате JSON в поле 'additional'."
	registrationRequestBody      = "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"refresh\": \"string\",\n    \"token\": \"string\"\n  }\n}"
	registrationResponseBodyInfo = "При успешной регистрации возвращаются два поля: 'token' и 'refresh'. " +
		"Эти токены следует сохранить для последующего использования в заголовках запросов."

	authorizationTitle            = "Авторизация"
	authorizationDescription      = "Авторизация пользователя с использованием email и пароля. После успешной авторизации возвращаются токены."
	authorizationRequestPath      = "https://asenmmo.com/api/v1/devToolsClient/signIn"
	authorizationRequestHeader    = "Content-Type: application/json"
	authorizationResponseBody     = "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"params\": {\n    \"client\": {\n      \"additional\": {},\n      \"email\": \"string\",\n      \"gameID\": \"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\n      \"nickname\": \"string\",\n      \"password\": \"string\"\n    }\n  }\n}"
	authorizationRequestBodyInfo  = "Передайте email и пароль для авторизации. В случае успеха вернутся токены."
	authorizationRequestBody      = "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\",\n  \"result\": {\n    \"refresh\": \"string\",\n    \"token\": \"string\"\n  }\n}"
	authorizationResponseBodyInfo = "Возвращаются 'token' и 'refresh', которые необходимо использовать для дальнейших запросов."
	authorizationErrorBody        = "{\n  \"id\": 1,\n  \"jsonrpc\": \"2.0\"}"

	refreshTokenTitle = "Обновление токенов"
	refreshTokenInfo  = "Данный метод используется для обновления токенов доступа и обновления. " +
		"Необходимо передать текущие токены в заголовках запроса. В теле запроса нет дополнительных параметров."
	refreshTokenRequestPath   = "/api/v1/devToolsClient/refreshToken"
	refreshTokenRequestHeader = "Token: ваш токен\nRefreshToken: ваш refresh токен"
	refreshTokenRequestBody   = `{
	"id": 1,
	"jsonrpc": "2.0",
	"params": {}
}`
	refreshTokenResponseBody = `{
	"id": 1,
	"jsonrpc": "2.0",
	"result": {
		"newRefresh": "string",
		"newToken": "string"
	}
}`
	refreshTokenResponseBodyInfo = "В ответе возвращаются новые токены: 'newRefresh' и 'newToken'. " +
		"Эти токены следует сохранить и использовать для дальнейших запросов."

	clientDataTitle = "Получение данных клиента"
	clientDataInfo  = "Этот метод используется для получения информации о клиенте на основе идентификатора игры (gameID). " +
		"Запрос требует наличие токена в заголовке для авторизации. В ответе возвращаются данные клиента, такие как email, nickname, gameID, а также дополнительная информация."
	clientDataRequestPath   = "/api/v1/devToolsClient/getClient"
	clientDataRequestHeader = "Token: ваш токен"
	clientDataRequestBody   = `{
	"id": 1,
	"jsonrpc": "2.0",
	"params": {
		"gameID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	}
}`
	clientDataRequestBodyInfo = "В теле запроса передается идентификатор игры (gameID). Необходимо указать токен в заголовке для авторизации запроса."
	clientDataResponseBody    = `{
	"id": 1,
	"jsonrpc": "2.0",
	"result": {
		"client": {
			"additional": {},
			"email": "string",
			"gameID": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			"nickname": "string"
		}
	}
}`
	clientDataResponseBodyInfo = "В ответе возвращается объект клиента, который содержит email, gameID, nickname и дополнительную информацию в поле 'additional'."

	updateClientTitle = "Обновление данных клиента"
	updateClientInfo  = "Данный метод используется для обновления информации о клиенте. " +
		"Необходимо передать токен в заголовке запроса и объект клиента в теле запроса. " +
		"Объект клиента может включать новый пароль, если он требуется."
	updateClientRequestPath   = "/api/v1/devToolsClient/updateClient"
	updateClientRequestHeader = "Token: ваш токен"
	updateClientRequestBody   = `{
	"id": 1,
	"jsonrpc": "2.0",
	"params": {
		"client": {
			"additional": {},
			"email": "string",
			"nickname": "string",
			"password": "string",
			"newPassword": "string"
		}
	}
}`
	updateClientRequestBodyInfo = "В теле запроса передается объект клиента, который включает поля: \n" +
		"- `additional`: дополнительная информация (может быть пустым объектом)\n" +
		"- `email`: новый email клиента\n" +
		"- `nickname`: новый nickname клиента\n" +
		"- `password`: текущий пароль клиента (для проверки)\n" +
		"- `newPassword`: новый пароль клиента (если требуется изменить)."
	updateClientResponseBody = `{
	"id": 1,
	"jsonrpc": "2.0",
	"result": {}
}`
	updateClientResponseBodyInfo = "При успешном обновлении возвращается объект, содержащий поля:\n" +
		"- `id`: идентификатор запроса, который совпадает с тем, что был отправлен.\n" +
		"- `jsonrpc`: версия JSON-RPC, которая также совпадает с той, что была отправлена.\n" +
		"- `result`: пустой объект, подтверждающий успешное выполнение запроса. " +
		"Этот объект указывает на то, что данные клиента были успешно обновлены."
)
