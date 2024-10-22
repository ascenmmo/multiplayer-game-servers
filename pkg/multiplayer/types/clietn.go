package types

import "github.com/google/uuid"

type Client struct {
	ID uuid.UUID `json:"-" bson:"_id"`

	Email       string `json:"email" bson:"email"`
	Nickname    string `json:"nickname" bson:"nickname"`
	Password    string `json:"password,omitempty" bson:"password"`
	NewPassword string `json:"newPassword,omitempty" bson:"-"`

	GameID uuid.UUID `json:"gameID" bson:"gameID"`

	Additional interface{} `json:"additional" bson:"additional"`
}
