package types

import (
	"github.com/google/uuid"
)

type Room struct {
	ID     uuid.UUID `json:"id" bson:"_id"`
	GameID uuid.UUID `json:"gameID" bson:"gameID"`
	Name   string    `json:"name" bson:"name"`

	CreatorID uuid.UUID `json:"creator_id" bson:"creator_id"`

	Connections []uuid.UUID `json:"connections" bson:"connections"`
	Servers     []uuid.UUID `json:"servers" bson:"servers"`

	IsExists  bool  `json:"is_exists" bson:"is_exists"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
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
