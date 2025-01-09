package types

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	ID     uuid.UUID `json:"id" bson:"_id"`
	GameID uuid.UUID `json:"gameID" bson:"gameID"`
	Name   string    `json:"name" bson:"name"`

	RoomCode string `json:"roomCode" bson:"roomCode"`

	CreatorID uuid.UUID `json:"creatorID" bson:"creatorID"`

	Connections   []uuid.UUID `json:"connections" bson:"connections"`
	Servers       []uuid.UUID `json:"servers" bson:"servers"`
	ExistsServers []uuid.UUID `json:"existsServers" bson:"existsServers"`

	IsExists  bool      `json:"isExists" bson:"isExists"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

func (r *Room) RemoveConnectionID(id uuid.UUID) {
	for i, conn := range r.Connections {
		if conn == id {
			r.Connections = append(r.Connections[:i], r.Connections[i+1:]...)
			break
		}
	}
}

func (r *Room) RemoveServerID(id uuid.UUID) {
	for i, conn := range r.Connections {
		if conn == id {
			r.Connections = append(r.Connections[:i], r.Connections[i+1:]...)
			break
		}
	}
}
