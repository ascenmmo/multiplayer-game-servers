package types

import "github.com/google/uuid"

type GameSaves struct {
	ID     uuid.UUID   `json:"-" bson:"_id"`
	GameID uuid.UUID   `json:"-" bson:"gameID"`
	UserID uuid.UUID   `json:"-" bson:"userID"`
	Saves  interface{} `json:"saves" bson:"saves"`
}
