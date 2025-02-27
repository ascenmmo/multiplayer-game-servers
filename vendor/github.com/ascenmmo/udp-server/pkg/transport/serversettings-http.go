// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/ascenmmo/udp-server/pkg/api"
	"github.com/gofiber/fiber/v2"
)

type httpServerSettings struct {
	errorHandler     ErrorHandler
	maxBatchSize     int
	maxParallelBatch int
	svc              *serverServerSettings
	base             api.ServerSettings
}

func NewServerSettings(svcServerSettings api.ServerSettings) (srv *httpServerSettings) {

	srv = &httpServerSettings{
		base: svcServerSettings,
		svc:  newServerServerSettings(svcServerSettings),
	}
	return
}

func (http *httpServerSettings) Service() *serverServerSettings {
	return http.svc
}

func (http *httpServerSettings) WithLog() *httpServerSettings {
	http.svc.WithLog()
	return http
}

func (http *httpServerSettings) WithTrace() *httpServerSettings {
	http.svc.WithTrace()
	return http
}

func (http *httpServerSettings) WithErrorHandler(handler ErrorHandler) *httpServerSettings {
	http.errorHandler = handler
	return http
}

func (http *httpServerSettings) SetRoutes(route *fiber.App) {
	route.Post("/api/v1/udp/serverSettings", http.serveBatch)
	route.Post("/api/v1/udp/serverSettings/getConnectionsNum", http.serveGetConnectionsNum)
	route.Post("/api/v1/udp/serverSettings/healthCheck", http.serveHealthCheck)
	route.Post("/api/v1/udp/serverSettings/getServerSettings", http.serveGetServerSettings)
	route.Post("/api/v1/udp/serverSettings/createRoom", http.serveCreateRoom)
	route.Post("/api/v1/udp/serverSettings/getDeletedRooms", http.serveGetDeletedRooms)
}
