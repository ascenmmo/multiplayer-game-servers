// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/gofiber/fiber/v2"
)

type httpDevToolsConnections struct {
	errorHandler     ErrorHandler
	maxBatchSize     int
	maxParallelBatch int
	svc              *serverDevToolsConnections
	base             multiplayer.DevToolsConnections
}

func NewDevToolsConnections(svcDevToolsConnections multiplayer.DevToolsConnections) (srv *httpDevToolsConnections) {

	srv = &httpDevToolsConnections{
		base: svcDevToolsConnections,
		svc:  newServerDevToolsConnections(svcDevToolsConnections),
	}
	return
}

func (http *httpDevToolsConnections) Service() *serverDevToolsConnections {
	return http.svc
}

func (http *httpDevToolsConnections) WithLog() *httpDevToolsConnections {
	http.svc.WithLog()
	return http
}

func (http *httpDevToolsConnections) WithTrace() *httpDevToolsConnections {
	http.svc.WithTrace()
	return http
}

func (http *httpDevToolsConnections) WithMetrics() *httpDevToolsConnections {
	http.svc.WithMetrics()
	return http
}

func (http *httpDevToolsConnections) WithErrorHandler(handler ErrorHandler) *httpDevToolsConnections {
	http.errorHandler = handler
	return http
}

func (http *httpDevToolsConnections) SetRoutes(route *fiber.App) {
	route.Post("/api/v1/devToolsConnections", http.serveBatch)
	route.Post("/api/v1/devToolsConnections/createRoom", http.serveCreateRoom)
	route.Post("/api/v1/devToolsConnections/getRoomsAll", http.serveGetRoomsAll)
	route.Post("/api/v1/devToolsConnections/joinRoomByID", http.serveJoinRoomByID)
	route.Post("/api/v1/devToolsConnections/removeRoomByID", http.serveRemoveRoomByID)
	route.Post("/api/v1/devToolsConnections/getRoomsConnectionUrls", http.serveGetRoomsConnectionUrls)
}
