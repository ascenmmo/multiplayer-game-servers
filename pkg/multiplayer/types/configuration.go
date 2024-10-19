package types

type ServerConfiguration struct {
	Address                       string                          `json:"address"`
	ServerPort                    string                          `json:"server_port"`
	MongoConnection               string                          `json:"mongo_connection"`
	TokenKey                      string                          `json:"token_key"`
	AdditionalServerConfiguration []AdditionalServerConfiguration `json:"additional_server_configuration"`
}

type AdditionalServerConfiguration struct {
	ServerType     string `json:"server_type"`
	ServerRestPort string `json:"server_rest_port"`
	ServerWSPort   string `json:"server_ws_port"`
	ServerUDPPort  string `json:"server_udp_port"`
}
