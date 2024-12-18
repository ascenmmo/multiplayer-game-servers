// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"fmt"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/go-kit/kit/metrics"
	"github.com/google/uuid"
	"time"
)

type metricsDevToolsClient struct {
	next            multiplayer.DevToolsClient
	requestCount    metrics.Counter
	requestCountAll metrics.Counter
	requestLatency  metrics.Histogram
}

func metricsMiddlewareDevToolsClient(next multiplayer.DevToolsClient) multiplayer.DevToolsClient {
	return &metricsDevToolsClient{
		next:            next,
		requestCount:    RequestCount.With("service", "DevToolsClient"),
		requestCountAll: RequestCountAll.With("service", "DevToolsClient"),
		requestLatency:  RequestLatency.With("service", "DevToolsClient"),
	}
}

func (m metricsDevToolsClient) SignUp(ctx context.Context, client types.Client) (token string, refresh string, err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "signUp", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "signUp", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "signUp").Add(1)

	return m.next.SignUp(ctx, client)
}

func (m metricsDevToolsClient) SignIn(ctx context.Context, client types.Client) (token string, refresh string, err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "signIn", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "signIn", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "signIn").Add(1)

	return m.next.SignIn(ctx, client)
}

func (m metricsDevToolsClient) RefreshToken(ctx context.Context, token string, refresh string) (newToken string, newRefresh string, err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "refreshToken", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "refreshToken", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "refreshToken").Add(1)

	return m.next.RefreshToken(ctx, token, refresh)
}

func (m metricsDevToolsClient) GetClient(ctx context.Context, token string, gameID uuid.UUID) (client types.Client, err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "getClient", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "getClient", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "getClient").Add(1)

	return m.next.GetClient(ctx, token, gameID)
}

func (m metricsDevToolsClient) UpdateClient(ctx context.Context, token string, client types.Client) (err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "updateClient", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "updateClient", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "updateClient").Add(1)

	return m.next.UpdateClient(ctx, token, client)
}

func (m metricsDevToolsClient) GetGameSaves(ctx context.Context, token string) (gameSaves types.GameSaves, err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "getGameSaves", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "getGameSaves", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "getGameSaves").Add(1)

	return m.next.GetGameSaves(ctx, token)
}

func (m metricsDevToolsClient) SetGameSaves(ctx context.Context, token string, gameSaves types.GameSaves) (err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "setGameSaves", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "setGameSaves", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "setGameSaves").Add(1)

	return m.next.SetGameSaves(ctx, token, gameSaves)
}

func (m metricsDevToolsClient) DeleteGameSaves(ctx context.Context, token string) (err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "deleteGameSaves", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "deleteGameSaves", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "deleteGameSaves").Add(1)

	return m.next.DeleteGameSaves(ctx, token)
}
