package types

import "github.com/google/uuid"

type RequestSetMessage struct {
	Server *uuid.UUID `json:"server,omitempty"`
	Token  string     `json:"token,omitempty"`
	Data   any        `json:"data"`
}

type ResponseGetMessage struct {
	DataArray []any `json:"dataArray"`
}

type CreateRoomRequest struct {
	TTL         string      `json:"time_to_live"`
	GameConfigs GameConfigs `json:"gameConfigs"`
}
