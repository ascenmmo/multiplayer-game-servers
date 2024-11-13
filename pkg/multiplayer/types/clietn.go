package types

import (
	"github.com/google/uuid"
	"strings"
)

type Client struct {
	ID uuid.UUID `json:"-" bson:"_id"`

	Email       string `json:"email" bson:"email"`
	Nickname    string `json:"nickname" bson:"nickname"`
	Password    string `json:"password,omitempty" bson:"password"`
	NewPassword string `json:"newPassword,omitempty" bson:"-"`

	GameID uuid.UUID `json:"gameID" bson:"gameID"`

	Additional interface{} `json:"additional" bson:"additional"`
}

func (c *Client) Update(newClient Client) {
	if newClient.Email != "" && c.Email != newClient.Email {
		c.Email = strings.ToLower(newClient.Email)

		c.ID = uuid.NewMD5(uuid.NameSpaceX500, []byte(newClient.GameID.String()+newClient.Email+newClient.Nickname))
	}
	if newClient.Nickname != "" && c.Nickname != newClient.Nickname {
		c.Nickname = newClient.Nickname
	}
	if newClient.Additional != "" && c.Additional != newClient.Additional {
		c.Additional = newClient.Additional
	}
}
