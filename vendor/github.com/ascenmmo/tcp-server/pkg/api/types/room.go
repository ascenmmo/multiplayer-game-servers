package types

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	GameID uuid.UUID
	RoomID uuid.UUID

	ServerID []uuid.UUID

	Users []*User

	UpdatedAt time.Time
}

type User struct {
	ID   uuid.UUID
	Data []interface{}
}

func (r *Room) SetUser(user *User) {
	r.Users = r.setUser(r.Users, user)
}

func (r *Room) GetUsers() (users []*User) {
	users = r.Users
	return
}

func (r *Room) RemoveUser(user uuid.UUID) {
	r.removeFromArray(user)
}

func (r *Room) SetUpdatedAt() {
	r.UpdatedAt = time.Now()
}

func (r *Room) setUser(users []*User, user *User) (allUsers []*User) {
	uniqueUsers := make(map[uuid.UUID]*User)
	for _, user := range users {
		uniqueUsers[user.ID] = user
	}
	uniqueUsers[user.ID] = user

	for _, user := range uniqueUsers {
		allUsers = append(allUsers, user)
	}

	return allUsers
}

func (r *Room) removeFromArray(userID uuid.UUID) {
	for i, user := range r.Users {
		if user.ID == userID {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
		}
	}
}

func (r *Room) SetServerID(id uuid.UUID) {
	for _, serverID := range r.ServerID {
		if serverID == id {
			return
		}
	}
	r.ServerID = append(r.ServerID, id)
}

func (r *Room) RemoveServerID(id uuid.UUID) {
	for i, server := range r.ServerID {
		if server == id {
			r.ServerID = append(r.ServerID[:i], r.ServerID[i+1:]...)
		}
	}
}
