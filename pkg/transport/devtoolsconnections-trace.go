// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

type traceDevToolsConnections struct {
	next multiplayer.DevToolsConnections
}

func traceMiddlewareDevToolsConnections(next multiplayer.DevToolsConnections) multiplayer.DevToolsConnections {
	return &traceDevToolsConnections{next: next}
}

func (svc traceDevToolsConnections) CreateRoom(ctx context.Context, token string, name string) (newToken string, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "CreateRoom")
	return svc.next.CreateRoom(ctx, token, name)
}

func (svc traceDevToolsConnections) GetRoomsAll(ctx context.Context, token string) (rooms []types.Room, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "GetRoomsAll")
	return svc.next.GetRoomsAll(ctx, token)
}

func (svc traceDevToolsConnections) JoinRoomByID(ctx context.Context, token string, roomID uuid.UUID) (newToken string, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "JoinRoomByID")
	return svc.next.JoinRoomByID(ctx, token, roomID)
}

func (svc traceDevToolsConnections) JoinRoomByRoomCode(ctx context.Context, token string, roomCode string) (newToken string, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "JoinRoomByRoomCode")
	return svc.next.JoinRoomByRoomCode(ctx, token, roomCode)
}

func (svc traceDevToolsConnections) GetMyRoom(ctx context.Context, token string) (room types.Room, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "GetMyRoom")
	return svc.next.GetMyRoom(ctx, token)
}

func (svc traceDevToolsConnections) LeaveRoom(ctx context.Context, token string, roomID uuid.UUID) (err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "LeaveRoom")
	return svc.next.LeaveRoom(ctx, token, roomID)
}

func (svc traceDevToolsConnections) RemoveRoomByID(ctx context.Context, token string, roomID uuid.UUID) (err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "RemoveRoomByID")
	return svc.next.RemoveRoomByID(ctx, token, roomID)
}

func (svc traceDevToolsConnections) GetRoomsConnectionUrls(ctx context.Context, token string) (connectionsServer []types.ConnectionServer, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "GetRoomsConnectionUrls")
	return svc.next.GetRoomsConnectionUrls(ctx, token)
}
