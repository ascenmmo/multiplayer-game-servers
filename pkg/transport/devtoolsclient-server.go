// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type serverDevToolsClient struct {
	svc             multiplayer.DevToolsClient
	signUp          DevToolsClientSignUp
	signIn          DevToolsClientSignIn
	refreshToken    DevToolsClientRefreshToken
	getClient       DevToolsClientGetClient
	updateClient    DevToolsClientUpdateClient
	getGameSaves    DevToolsClientGetGameSaves
	setGameSaves    DevToolsClientSetGameSaves
	deleteGameSaves DevToolsClientDeleteGameSaves
}

type MiddlewareSetDevToolsClient interface {
	Wrap(m MiddlewareDevToolsClient)
	WrapSignUp(m MiddlewareDevToolsClientSignUp)
	WrapSignIn(m MiddlewareDevToolsClientSignIn)
	WrapRefreshToken(m MiddlewareDevToolsClientRefreshToken)
	WrapGetClient(m MiddlewareDevToolsClientGetClient)
	WrapUpdateClient(m MiddlewareDevToolsClientUpdateClient)
	WrapGetGameSaves(m MiddlewareDevToolsClientGetGameSaves)
	WrapSetGameSaves(m MiddlewareDevToolsClientSetGameSaves)
	WrapDeleteGameSaves(m MiddlewareDevToolsClientDeleteGameSaves)

	WithTrace()
	WithMetrics()
	WithLog()
}

func newServerDevToolsClient(svc multiplayer.DevToolsClient) *serverDevToolsClient {
	return &serverDevToolsClient{
		deleteGameSaves: svc.DeleteGameSaves,
		getClient:       svc.GetClient,
		getGameSaves:    svc.GetGameSaves,
		refreshToken:    svc.RefreshToken,
		setGameSaves:    svc.SetGameSaves,
		signIn:          svc.SignIn,
		signUp:          svc.SignUp,
		svc:             svc,
		updateClient:    svc.UpdateClient,
	}
}

func (srv *serverDevToolsClient) Wrap(m MiddlewareDevToolsClient) {
	srv.svc = m(srv.svc)
	srv.signUp = srv.svc.SignUp
	srv.signIn = srv.svc.SignIn
	srv.refreshToken = srv.svc.RefreshToken
	srv.getClient = srv.svc.GetClient
	srv.updateClient = srv.svc.UpdateClient
	srv.getGameSaves = srv.svc.GetGameSaves
	srv.setGameSaves = srv.svc.SetGameSaves
	srv.deleteGameSaves = srv.svc.DeleteGameSaves
}

func (srv *serverDevToolsClient) SignUp(ctx context.Context, client types.Client) (token string, refresh string, err error) {
	return srv.signUp(ctx, client)
}

func (srv *serverDevToolsClient) SignIn(ctx context.Context, client types.Client) (token string, refresh string, err error) {
	return srv.signIn(ctx, client)
}

func (srv *serverDevToolsClient) RefreshToken(ctx context.Context, token string, refresh string) (newToken string, newRefresh string, err error) {
	return srv.refreshToken(ctx, token, refresh)
}

func (srv *serverDevToolsClient) GetClient(ctx context.Context, token string, gameID uuid.UUID) (client types.Client, err error) {
	return srv.getClient(ctx, token, gameID)
}

func (srv *serverDevToolsClient) UpdateClient(ctx context.Context, token string, client types.Client) (err error) {
	return srv.updateClient(ctx, token, client)
}

func (srv *serverDevToolsClient) GetGameSaves(ctx context.Context, token string) (gameSaves types.GameSaves, err error) {
	return srv.getGameSaves(ctx, token)
}

func (srv *serverDevToolsClient) SetGameSaves(ctx context.Context, token string, gameSaves types.GameSaves) (err error) {
	return srv.setGameSaves(ctx, token, gameSaves)
}

func (srv *serverDevToolsClient) DeleteGameSaves(ctx context.Context, token string) (err error) {
	return srv.deleteGameSaves(ctx, token)
}

func (srv *serverDevToolsClient) WrapSignUp(m MiddlewareDevToolsClientSignUp) {
	srv.signUp = m(srv.signUp)
}

func (srv *serverDevToolsClient) WrapSignIn(m MiddlewareDevToolsClientSignIn) {
	srv.signIn = m(srv.signIn)
}

func (srv *serverDevToolsClient) WrapRefreshToken(m MiddlewareDevToolsClientRefreshToken) {
	srv.refreshToken = m(srv.refreshToken)
}

func (srv *serverDevToolsClient) WrapGetClient(m MiddlewareDevToolsClientGetClient) {
	srv.getClient = m(srv.getClient)
}

func (srv *serverDevToolsClient) WrapUpdateClient(m MiddlewareDevToolsClientUpdateClient) {
	srv.updateClient = m(srv.updateClient)
}

func (srv *serverDevToolsClient) WrapGetGameSaves(m MiddlewareDevToolsClientGetGameSaves) {
	srv.getGameSaves = m(srv.getGameSaves)
}

func (srv *serverDevToolsClient) WrapSetGameSaves(m MiddlewareDevToolsClientSetGameSaves) {
	srv.setGameSaves = m(srv.setGameSaves)
}

func (srv *serverDevToolsClient) WrapDeleteGameSaves(m MiddlewareDevToolsClientDeleteGameSaves) {
	srv.deleteGameSaves = m(srv.deleteGameSaves)
}

func (srv *serverDevToolsClient) WithTrace() {
	srv.Wrap(traceMiddlewareDevToolsClient)
}

func (srv *serverDevToolsClient) WithMetrics() {
	srv.Wrap(metricsMiddlewareDevToolsClient)
}

func (srv *serverDevToolsClient) WithLog() {
	srv.Wrap(loggerMiddlewareDevToolsClient())
}
