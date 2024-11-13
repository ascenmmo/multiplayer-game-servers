package env

var (
	ServerAddress       = "0.0.0.0"                          // Server IP address
	TCPPort             = "8083"                             // Port for TCP connections
	TokenKey            = "_remember_token_must_be_32_bytes" // Unique token for authentication
	MaxRequestPerSecond = 5                                  // Max requests per second
)
