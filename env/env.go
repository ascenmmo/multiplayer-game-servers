package env

var (
	ServerAddress = "ascenmmo.com"
	TokenKey      = "_remember_token_mast_be_32_bytes"
	DebugLogs     = true
)

var (
	RunMultiplayer                 = true
	RunAdminPanel                  = true
	MultiplayerPort                = "8080"
	MongoURL                       = "mongodb://username:userpassword@ascenmmo.com:27017"
	MultiplayerMaxRequestPerSecond = 100
)

var (
	RunUdpServer                 = true
	UdpServerPort                = "8081"
	UdpServerConnectionPort      = "4500"
	UdpServerMaxRequestPerSecond = 200
)

var (
	RunWebsocketServer                 = true
	WebsocketServerPort                = "8082"
	WebsocketServerConnectionPort      = "4240"
	WebsocketServerMaxRequestPerSecond = 100
)

var (
	RunTcpServer                 = true
	TcpServerPort                = "8083"
	TcpServerMaxRequestPerSecond = 5
)
