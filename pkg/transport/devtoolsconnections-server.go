// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type serverDevToolsConnections struct {
	svc                    multiplayer.DevToolsConnections
	createRoom             DevToolsConnectionsCreateRoom
	getRoomsAll            DevToolsConnectionsGetRoomsAll
	joinRoomByID           DevToolsConnectionsJoinRoomByID
	removeRoomByID         DevToolsConnectionsRemoveRoomByID
	getRoomsConnectionUrls DevToolsConnectionsGetRoomsConnectionUrls
}

type MiddlewareSetDevToolsConnections interface {
	Wrap(m MiddlewareDevToolsConnections)
	WrapCreateRoom(m MiddlewareDevToolsConnectionsCreateRoom)
	WrapGetRoomsAll(m MiddlewareDevToolsConnectionsGetRoomsAll)
	WrapJoinRoomByID(m MiddlewareDevToolsConnectionsJoinRoomByID)
	WrapRemoveRoomByID(m MiddlewareDevToolsConnectionsRemoveRoomByID)
	WrapGetRoomsConnectionUrls(m MiddlewareDevToolsConnectionsGetRoomsConnectionUrls)

	WithTrace()
	WithMetrics()
	WithLog()
}

func newServerDevToolsConnections(svc multiplayer.DevToolsConnections) *serverDevToolsConnections {
	return &serverDevToolsConnections{
		createRoom:             svc.CreateRoom,
		getRoomsAll:            svc.GetRoomsAll,
		getRoomsConnectionUrls: svc.GetRoomsConnectionUrls,
		joinRoomByID:           svc.JoinRoomByID,
		removeRoomByID:         svc.RemoveRoomByID,
		svc:                    svc,
	}
}

func (srv *serverDevToolsConnections) Wrap(m MiddlewareDevToolsConnections) {
	srv.svc = m(srv.svc)
	srv.createRoom = srv.svc.CreateRoom
	srv.getRoomsAll = srv.svc.GetRoomsAll
	srv.joinRoomByID = srv.svc.JoinRoomByID
	srv.removeRoomByID = srv.svc.RemoveRoomByID
	srv.getRoomsConnectionUrls = srv.svc.GetRoomsConnectionUrls
}

func (srv *serverDevToolsConnections) CreateRoom(ctx context.Context, token string, gameID uuid.UUID) (newToken string, err error) {
	return srv.createRoom(ctx, token, gameID)
}

func (srv *serverDevToolsConnections) GetRoomsAll(ctx context.Context, token string, gameID uuid.UUID) (rooms []types.Room, err error) {
	return srv.getRoomsAll(ctx, token, gameID)
}

func (srv *serverDevToolsConnections) JoinRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (newToken string, err error) {
	return srv.joinRoomByID(ctx, token, gameID, roomID)
}

func (srv *serverDevToolsConnections) RemoveRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (err error) {
	return srv.removeRoomByID(ctx, token, gameID, roomID)
}

func (srv *serverDevToolsConnections) GetRoomsConnectionUrls(ctx context.Context, token string) (connectionsServer []types.ConnectionServer, err error) {
	return srv.getRoomsConnectionUrls(ctx, token)
}

func (srv *serverDevToolsConnections) WrapCreateRoom(m MiddlewareDevToolsConnectionsCreateRoom) {
	srv.createRoom = m(srv.createRoom)
}

func (srv *serverDevToolsConnections) WrapGetRoomsAll(m MiddlewareDevToolsConnectionsGetRoomsAll) {
	srv.getRoomsAll = m(srv.getRoomsAll)
}

func (srv *serverDevToolsConnections) WrapJoinRoomByID(m MiddlewareDevToolsConnectionsJoinRoomByID) {
	srv.joinRoomByID = m(srv.joinRoomByID)
}

func (srv *serverDevToolsConnections) WrapRemoveRoomByID(m MiddlewareDevToolsConnectionsRemoveRoomByID) {
	srv.removeRoomByID = m(srv.removeRoomByID)
}

func (srv *serverDevToolsConnections) WrapGetRoomsConnectionUrls(m MiddlewareDevToolsConnectionsGetRoomsConnectionUrls) {
	srv.getRoomsConnectionUrls = m(srv.getRoomsConnectionUrls)
}

func (srv *serverDevToolsConnections) WithTrace() {
	srv.Wrap(traceMiddlewareDevToolsConnections)
}

func (srv *serverDevToolsConnections) WithMetrics() {
	srv.Wrap(metricsMiddlewareDevToolsConnections)
}

func (srv *serverDevToolsConnections) WithLog() {
	srv.Wrap(loggerMiddlewareDevToolsConnections())
}
