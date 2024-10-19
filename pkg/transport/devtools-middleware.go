// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type DevToolsCreateGame func(ctx context.Context, token string, game types.Game) (id uuid.UUID, err error)
type DevToolsGameAddOwner func(ctx context.Context, token string, gameID uuid.UUID, userID uuid.UUID) (err error)
type DevToolsGameRemoveUser func(ctx context.Context, token string, gameID uuid.UUID, userID uuid.UUID) (err error)
type DevToolsUpdateGame func(ctx context.Context, token string, gameID uuid.UUID, newGame types.Game) (id uuid.UUID, err error)
type DevToolsGetMyGames func(ctx context.Context, token string) (games []types.Game, err error)
type DevToolsGetGameByGameID func(ctx context.Context, token string, gameID uuid.UUID) (game types.Game, err error)
type DevToolsTurnOnServerInGame func(ctx context.Context, token string, serverID uuid.UUID, gameId uuid.UUID) (err error)
type DevToolsTurnOffServerInGame func(ctx context.Context, token string, serverID uuid.UUID, gameId uuid.UUID) (err error)

type MiddlewareDevTools func(next multiplayer.DevTools) multiplayer.DevTools

type MiddlewareDevToolsCreateGame func(next DevToolsCreateGame) DevToolsCreateGame
type MiddlewareDevToolsGameAddOwner func(next DevToolsGameAddOwner) DevToolsGameAddOwner
type MiddlewareDevToolsGameRemoveUser func(next DevToolsGameRemoveUser) DevToolsGameRemoveUser
type MiddlewareDevToolsUpdateGame func(next DevToolsUpdateGame) DevToolsUpdateGame
type MiddlewareDevToolsGetMyGames func(next DevToolsGetMyGames) DevToolsGetMyGames
type MiddlewareDevToolsGetGameByGameID func(next DevToolsGetGameByGameID) DevToolsGetGameByGameID
type MiddlewareDevToolsTurnOnServerInGame func(next DevToolsTurnOnServerInGame) DevToolsTurnOnServerInGame
type MiddlewareDevToolsTurnOffServerInGame func(next DevToolsTurnOffServerInGame) DevToolsTurnOffServerInGame
