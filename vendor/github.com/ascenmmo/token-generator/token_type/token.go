package tokentype

import (
	"github.com/google/uuid"
	"time"
)

type Info struct {
	GameID uuid.UUID     `json:"game_id" bson:"game_id"`
	RoomID uuid.UUID     `json:"room_id" bson:"room_id"`
	UserID uuid.UUID     `json:"user_id" bson:"user_id"`
	TTL    time.Duration `json:"ttl" bson:"ttl"`
}

func (i *Info) SetDefaultValueIfEmpty() {
	if i.TTL == 0 {
		i.TTL = time.Hour
	}
}
