package env

var (
	ServerAddress       = "127.0.0.1"                        // Server address
	TCPPort             = "8081"                             // Port for TCP connections
	UDPPort             = "4500"                             // Port for UDP connections
	TokenKey            = "_remember_token_must_be_32_bytes" // Unique token for authentication
	MaxRequestPerSecond = 200                                // Maximum number of requests per second
)
