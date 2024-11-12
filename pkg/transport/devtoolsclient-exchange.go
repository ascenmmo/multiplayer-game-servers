// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type requestDevToolsClientSignUp struct {
	Client types.Client `json:"client"`
}

type responseDevToolsClientSignUp struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

type requestDevToolsClientSignIn struct {
	Client types.Client `json:"client"`
}

type responseDevToolsClientSignIn struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

type requestDevToolsClientRefreshToken struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

type responseDevToolsClientRefreshToken struct {
	NewToken   string `json:"newToken"`
	NewRefresh string `json:"newRefresh"`
}

type requestDevToolsClientGetClient struct {
	Token  string    `json:"token"`
	GameID uuid.UUID `json:"gameID"`
}

type responseDevToolsClientGetClient struct {
	Client types.Client `json:"client"`
}

type requestDevToolsClientUpdateClient struct {
	Token  string       `json:"token"`
	Client types.Client `json:"client"`
}

// Formal exchange type, please do not delete.
type responseDevToolsClientUpdateClient struct{}

type requestDevToolsClientGetGameSaves struct {
	Token string `json:"token"`
}

type responseDevToolsClientGetGameSaves struct {
	GameSaves types.GameSaves `json:"gameSaves"`
}

type requestDevToolsClientSetGameSaves struct {
	Token     string          `json:"token"`
	GameSaves types.GameSaves `json:"gameSaves"`
}

// Formal exchange type, please do not delete.
type responseDevToolsClientSetGameSaves struct{}

type requestDevToolsClientDeleteGameSaves struct {
	Token string `json:"token"`
}

// Formal exchange type, please do not delete.
type responseDevToolsClientDeleteGameSaves struct{}
