package types

import (
	"github.com/google/uuid"
)

type AccessGame struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	GameID    uuid.UUID `json:"gameID" bson:"gameID"`
	CreatorID uuid.UUID `json:"creator_id" bson:"creator_id"`

	Owners []uuid.UUID `json:"owners" bson:"owners"`
}

type AccessClient struct {
	ID     uuid.UUID `json:"id" bson:"_id"`
	UserID uuid.UUID `json:"user_id" bson:"user_id"`
}

func (g *AccessGame) RemoveOwner(id uuid.UUID) {
	for i, ownersID := range g.Owners {
		if ownersID == id {
			g.Owners = append(g.Owners[:i], g.Owners[i+1:]...)
			break
		}
	}
}
