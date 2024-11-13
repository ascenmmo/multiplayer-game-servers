# ASCENMMO Admin Panel

## Русская версия

Для русскоязычных пользователей документацию можно найти [здесь](https://github.com/ascenmmo/multiplayer-game-servers/blob/master/RU_README.md).


## Description

The ASCENMMO admin panel provides an easy-to-use interface for managing game servers. You can register a new game, and the admin panel will return a unique game ID, which can then be used for registering players in the game.

### How to Use the Admin Panel

1. Go to the link [ascenmmo.com/admin/games](ascenmmo.com/admin/games) to access the admin panel.
2. Register a new game by filling out the required fields.
3. Obtain a unique game ID that will be used by clients to register players.

## Running Game Servers

With the admin panel, you can also launch other game servers, such as UDP, WebSocket, and TCP. This will allow you to create high-performance multiplayer games.

## Documentation

Detailed documentation for connecting is available at [ascenmmo.com/developer/doc](http://ascenmmo.com/developer/doc).

## Environment Variables

To ensure the service functions correctly, you need to set the environment variables. Here’s an example you can use:

```go
package env

var (
	ServerAddress = "ascenmmo.com" // Server address
	TokenKey      = "_remember_token_must_be_32_bytes" // Unique authentication token
)

var (
	RunMultiplayer                 = true // Enable multiplayer mode
	RunAdminPanel                  = true // Enable admin panel
	MultiplayerPort                = "8080" // Port for multiplayer mode
	MongoURL                       = "mongodb://username:userpassword@ascenmmo.com:27017" // MongoDB connection URL
	MultiplayerMaxRequestPerSecond = 5 // Maximum requests per second for multiplayer mode
)

var (
	RunUdpServer                 = true // Enable UDP server
	UdpServerPort                = "8081" // Port for UDP server
	UdpServerConnectionPort      = "4500" // Connection port for UDP
	UdpServerMaxRequestPerSecond = 200 // Maximum requests per second for UDP server
)

var (
	RunWebsocketServer                 = true // Enable WebSocket server
	WebsocketServerPort                = "8082" // Port for WebSocket server
	WebsocketServerConnectionPort      = "4240" // Connection port for WebSocket
	WebsocketServerMaxRequestPerSecond = 100 // Maximum requests per second for WebSocket server
)

var (
	RunTcpServer                 = true // Enable TCP server
	TcpServerPort                = "8083" // Port for TCP server
	TcpServerMaxRequestPerSecond = 5 // Maximum requests per second for TCP server
)

```

## Docker Setup

Make sure that the ports specified in the environment variables correspond to the settings in your `docker-compose.yml`. If you are using MongoDB, create a `db.env` file with the following variables to initialize the database:

```plaintext
MONGO_INITDB_ROOT_USERNAME=username
MONGO_INITDB_ROOT_PASSWORD=userpassword
```
Also, update the new `MONGO_INITDB_ROOT_USERNAME` and `MONGO_INITDB_ROOT_PASSWORD`` in the config file `env/env.go` under the `MongoURL` variable.

### Starting the Project

After configuring the environment variables and docker-compose.yml, you can start the project with the following command:

```bash
docker-compose up -d --force-recreate --build
``` 

This command will create and run all services in the background, and rebuild them if necessary. If you need to free up space, you can use the command to clean up unused images:

## Accessing the Admin Panel
Once the project is running, you can access the admin panel at:
```bash
http://localhost:8080/admin/games
``` 
Here, you will need to register and add a game. Upon successful registration, the admin panel will return a unique ID for your game, which can be used for connecting game clients and registering players.


## Important Links

- [Developer Documentation](http://ascenmmo.com/developer/doc): Here you will find detailed instructions on connecting to game servers, as well as API usage examples.
- [Admin Panel](http://ascenmmo.com/admin/games): Use this link to open the admin panel, register your game, and obtain a unique ID for client connections.
- [GitHub Repository](https://github.com/ascenmmo): The project's source code, including server and client components. You can download it and run it on your server.

## Conclusion

With the ASCENMMO admin panel and the ability to launch game servers, you can easily manage multiplayer games and ensure their stable operation. If you have any questions or issues, please refer to the documentation or reach out to the developer community for assistance.

We hope this project will be useful for your game development and help you create engaging multiplayer experiences!



multiplayer, gaming, game-server, admin-panel, UDP, WebSocket, TCP, Go, Golang, MongoDB, docker, game-development, real-time, api, server-client, cross-platform, game-management, developer-tools, open-source, ASCENMMO


## Теги

`UDP`, `игровой сервер`, `высоконагруженный`, `бесплатное развертывание`, `Docker`, `кроссплатформенный`, `игровая разработка`, `сеть`, `многопользовательская игра`, `админка`, `управление играми`, `аутентификация`, `токены`, `Golang`, `open-source`, `MongoDB`, `WebSocket`, `TCP`, `API`, `реальное время`, `инструменты для разработчиков`

## Tags

`UDP`, `game server`, `high-performance`, `free deployment`, `Docker`, `cross-platform`, `game development`, `network`, `multiplayer game`, `admin panel`, `game management`, `authentication`, `tokens`, `Golang`, `open-source`, `MongoDB`, `WebSocket`, `TCP`, `API`, `real-time`, `developer tools`

