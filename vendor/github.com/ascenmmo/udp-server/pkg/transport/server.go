// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
)

type Server struct {
	log zerolog.Logger

	httpAfter  []Handler
	httpBefore []Handler

	config fiber.Config

	srvHTTP   *fiber.App
	srvHealth *fiber.App

	reporterCloser io.Closer

	maxBatchSize     int
	maxParallelBatch int

	httpServerSettings *httpServerSettings
	headerHandlers     map[string]HeaderHandler
}

func New(log zerolog.Logger, options ...Option) (srv *Server) {

	srv = &Server{
		config:           fiber.Config{DisableStartupMessage: true},
		headerHandlers:   make(map[string]HeaderHandler),
		log:              log,
		maxBatchSize:     defaultMaxBatchSize,
		maxParallelBatch: defaultMaxParallelBatch,
	}
	for _, option := range options {
		option(srv)
	}
	srv.srvHTTP = fiber.New(srv.config)
	srv.srvHTTP.Use(recoverHandler)
	srv.srvHTTP.Use(srv.setLogger)
	srv.srvHTTP.Use(srv.logLevelHandler)
	srv.srvHTTP.Use(srv.headersHandler)
	for _, option := range options {
		option(srv)
	}
	srv.srvHTTP.Post("/", srv.serveBatch)
	return
}

func (srv *Server) Fiber() *fiber.App {
	return srv.srvHTTP
}

func (srv *Server) WithLog() *Server {
	if srv.httpServerSettings != nil {
		srv.httpServerSettings = srv.ServerSettings().WithLog()
	}
	return srv
}

func (srv *Server) ServeHealth(address string, response interface{}) {
	srv.srvHealth = fiber.New(fiber.Config{DisableStartupMessage: true})
	srv.srvHealth.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(response)
	})
	go func() {
		err := srv.srvHealth.Listen(address)
		ExitOnError(srv.log, err, "serve health on "+address)
	}()
}

func sendResponse(ctx *fiber.Ctx, resp interface{}) (err error) {
	ctx.Response().Header.SetContentType("application/json")
	if err = json.NewEncoder(ctx).Encode(resp); err != nil {
		log.Ctx(ctx.UserContext()).Error().Err(err).Str("body", string(ctx.Body())).Msg("response write error")
	}
	return
}

func (srv *Server) Shutdown() {
	if srv.srvHTTP != nil {
		_ = srv.srvHTTP.Shutdown()
	}
	if srv.srvHealth != nil {
		_ = srv.srvHealth.Shutdown()
	}
}

func (srv *Server) WithTrace() *Server {
	if srv.httpServerSettings != nil {
		srv.httpServerSettings = srv.ServerSettings().WithTrace()
	}
	return srv
}

func (srv *Server) ServerSettings() *httpServerSettings {
	return srv.httpServerSettings
}
