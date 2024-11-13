// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type DevToolsClientSignUp func(ctx context.Context, client types.Client) (token string, refresh string, err error)
type DevToolsClientSignIn func(ctx context.Context, client types.Client) (token string, refresh string, err error)
type DevToolsClientRefreshToken func(ctx context.Context, token string, refresh string) (newToken string, newRefresh string, err error)
type DevToolsClientGetClient func(ctx context.Context, token string, gameID uuid.UUID) (client types.Client, err error)
type DevToolsClientUpdateClient func(ctx context.Context, token string, client types.Client) (err error)
type DevToolsClientGetGameSaves func(ctx context.Context, token string) (gameSaves types.GameSaves, err error)
type DevToolsClientSetGameSaves func(ctx context.Context, token string, gameSaves types.GameSaves) (err error)
type DevToolsClientDeleteGameSaves func(ctx context.Context, token string) (err error)

type MiddlewareDevToolsClient func(next multiplayer.DevToolsClient) multiplayer.DevToolsClient

type MiddlewareDevToolsClientSignUp func(next DevToolsClientSignUp) DevToolsClientSignUp
type MiddlewareDevToolsClientSignIn func(next DevToolsClientSignIn) DevToolsClientSignIn
type MiddlewareDevToolsClientRefreshToken func(next DevToolsClientRefreshToken) DevToolsClientRefreshToken
type MiddlewareDevToolsClientGetClient func(next DevToolsClientGetClient) DevToolsClientGetClient
type MiddlewareDevToolsClientUpdateClient func(next DevToolsClientUpdateClient) DevToolsClientUpdateClient
type MiddlewareDevToolsClientGetGameSaves func(next DevToolsClientGetGameSaves) DevToolsClientGetGameSaves
type MiddlewareDevToolsClientSetGameSaves func(next DevToolsClientSetGameSaves) DevToolsClientSetGameSaves
type MiddlewareDevToolsClientDeleteGameSaves func(next DevToolsClientDeleteGameSaves) DevToolsClientDeleteGameSaves
