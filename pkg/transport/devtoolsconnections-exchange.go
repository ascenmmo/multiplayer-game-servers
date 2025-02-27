// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type requestDevToolsConnectionsCreateRoom struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

type responseDevToolsConnectionsCreateRoom struct {
	NewToken string `json:"newToken"`
}

type requestDevToolsConnectionsGetRoomsAll struct {
	Token string `json:"token"`
}

type responseDevToolsConnectionsGetRoomsAll struct {
	Rooms []types.Room `json:"rooms"`
}

type requestDevToolsConnectionsJoinRoomByID struct {
	Token  string    `json:"token"`
	RoomID uuid.UUID `json:"roomID"`
}

type responseDevToolsConnectionsJoinRoomByID struct {
	NewToken string `json:"newToken"`
}

type requestDevToolsConnectionsJoinRoomByRoomCode struct {
	Token    string `json:"token"`
	RoomCode string `json:"roomCode"`
}

type responseDevToolsConnectionsJoinRoomByRoomCode struct {
	NewToken string `json:"newToken"`
}

type requestDevToolsConnectionsGetMyRoom struct {
	Token string `json:"token"`
}

type responseDevToolsConnectionsGetMyRoom struct {
	Room types.Room `json:"room"`
}

type requestDevToolsConnectionsLeaveRoom struct {
	Token  string    `json:"token"`
	RoomID uuid.UUID `json:"roomID"`
}

// Formal exchange type, please do not delete.
type responseDevToolsConnectionsLeaveRoom struct{}

type requestDevToolsConnectionsRemoveRoomByID struct {
	Token  string    `json:"token"`
	RoomID uuid.UUID `json:"roomID"`
}

// Formal exchange type, please do not delete.
type responseDevToolsConnectionsRemoveRoomByID struct{}

type requestDevToolsConnectionsGetRoomsConnectionUrls struct {
	Token string `json:"token"`
}

type responseDevToolsConnectionsGetRoomsConnectionUrls struct {
	ConnectionsServer []types.ConnectionServer `json:"connectionsServer"`
}
