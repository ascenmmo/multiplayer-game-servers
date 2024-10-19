package utils

import (
	"fmt"
	tokentype "github.com/ascenmmo/token-generator/token_type"
)

const (
	serverKey = "notify_server"
)

func GenerateRoomKey(clientInfo tokentype.Info) string {
	return fmt.Sprintf("game:%s-room:%s", clientInfo.GameID, clientInfo.RoomID)
}

func GenerateNotifyServerKey() string {
	return serverKey
}
