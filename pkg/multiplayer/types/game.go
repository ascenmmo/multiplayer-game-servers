package types

import "github.com/google/uuid"

type Game struct {
	ID          uuid.UUID `json:"gameID" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`

	CountPlayers int `json:"count_players" bson:"count_players"`
	CountOnline  int `json:"count_online" bson:"count_online"`

	Owners []uuid.UUID `json:"owners" bson:"owners"`

	Servers []uuid.UUID `json:"servers" bson:"servers"`
}

func (g *Game) RemoveOwners(id uuid.UUID) {
	for i, owner := range g.Owners {
		if owner == id {
			g.Owners = append(g.Owners[:i], g.Owners[i+1:]...)
			break
		}
	}
}

func (g *Game) RemoveServer(id uuid.UUID) {
	for i, server := range g.Servers {
		if server == id {
			g.Servers = append(g.Servers[:i], g.Servers[i+1:]...)
			break
		}
	}
}

func (g *Game) UpdateGame(newGame Game) (isUpdated bool) {
	if g.Name != newGame.Name && newGame.Name != "" {
		isUpdated = true
		g.Name = newGame.Name
	}

	if g.Description != newGame.Description && newGame.Description != "" {
		isUpdated = true
		g.Description = newGame.Description
	}
	return isUpdated
}

func (g *Game) RemoveEmptyArray() {
	if g.Owners == nil {
		g.Owners = []uuid.UUID{}
	}
	if g.Servers == nil {
		g.Servers = []uuid.UUID{}
	}
}
