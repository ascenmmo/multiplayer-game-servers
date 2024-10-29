package types

import "github.com/google/uuid"

type Developer struct {
	ID uuid.UUID `json:"-" bson:"_id"`

	Email       string `json:"email" bson:"email"`
	Nickname    string `json:"nickname" bson:"nickname"`
	Password    string `json:"password,omitempty" bson:"password"`
	NewPassword string `json:"new_password,omitempty" bson:"new_password"`
}
