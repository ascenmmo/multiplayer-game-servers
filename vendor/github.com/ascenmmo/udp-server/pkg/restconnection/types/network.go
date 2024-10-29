package types

import "github.com/google/uuid"

type Request struct {
	Server *uuid.UUID `json:"server,omitempty"`
	Token  string     `json:"token"`
	Data   any        `json:"data"`
}

type Response struct {
	Data any `json:"data"`
}

type CreateRoomRequest struct {
	GameConfigs GameConfigs `json:"game_configs"`
}
