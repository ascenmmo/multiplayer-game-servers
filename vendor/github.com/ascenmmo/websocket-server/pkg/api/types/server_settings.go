package types

import (
	"github.com/google/uuid"
	"time"
)

type CreateRoomRequest struct {
	RoomTTl time.Duration `json:"roomTTl"`
}

type GetDeletedRooms struct {
	GameID uuid.UUID `json:"gameID"`
	RoomID uuid.UUID `json:"roomID"`
}
