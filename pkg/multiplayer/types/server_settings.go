package types

import "github.com/google/uuid"

type GetDeletedRooms struct {
	GameID uuid.UUID `json:"gameID"`
	RoomID uuid.UUID `json:"roomID"`
}
