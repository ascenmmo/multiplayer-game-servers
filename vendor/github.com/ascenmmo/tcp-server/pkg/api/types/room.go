package types

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type Room struct {
	GameID uuid.UUID
	RoomID uuid.UUID

	ServerID []uuid.UUID

	Users []*User

	UpdatedAt time.Time
	mtx       sync.RWMutex
}

type User struct {
	ID   uuid.UUID
	Data []interface{}
}

func (r *Room) SetUser(user *User) {
	r.mtx.Lock()
	r.Users = r.setUser(r.Users, user)
	r.mtx.Unlock()
}

func (r *Room) GetUsers() (users []*User) {
	r.mtx.RLock()
	users = r.Users
	r.mtx.RUnlock()
	return
}

func (r *Room) RemoveUser(user uuid.UUID) {
	r.mtx.Lock()
	r.removeFromArray(user)
	r.mtx.Unlock()
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
	r.mtx.Lock()
	for i, user := range r.Users {
		if user.ID == userID {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
		}
	}
	r.mtx.Unlock()
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
	r.mtx.Lock()
	for i, server := range r.ServerID {
		if server == id {
			r.ServerID = append(r.ServerID[:i], r.ServerID[i+1:]...)
		}
	}
	r.mtx.Unlock()
}
