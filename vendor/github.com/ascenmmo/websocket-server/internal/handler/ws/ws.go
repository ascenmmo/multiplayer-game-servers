package wsconnection

import (
	"context"
	"fmt"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/websocket-server/internal/connection"
	"github.com/ascenmmo/websocket-server/internal/service"
	"github.com/ascenmmo/websocket-server/internal/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"net/http"
	"sync"
	"time"
)

const (
	bufferSize   = 32 * 1024
	readTimeout  = 60 * time.Second
	pingInterval = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  64 * 1024 * 100,
	WriteBufferSize: 64 * 1024 * 100,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocket struct {
	server    *http.Server
	mtx       sync.RWMutex
	service   service.Service
	rateLimit utils.RateLimit
	logger    zerolog.Logger
}

func (ws *WebSocket) connect(w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("token")
	if token == "" {
		http.Error(w, "token required", http.StatusBadRequest)
		return
	}

	clientInfo, err := ws.service.ParseToken(token)
	if err != nil {
		http.Error(w, "wrong token", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		ws.logger.Error().Err(err).Msg("failed to upgrade connection")
		return
	}

	ws.handleConnection(context.Background(), token, clientInfo, conn)
}

func (ws *WebSocket) handleConnection(ctx context.Context, token string, clientInfo tokentype.Info, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		err := ws.service.RemoveUser(clientInfo, clientInfo.UserID)
		if err != nil {
			ws.logger.Error().Err(err).Msg("failed to remove user")
		}
	}()

	conn.SetReadLimit(bufferSize)
	err := conn.SetReadDeadline(time.Now().Add(readTimeout))
	if err != nil {
		ws.logger.Error().Err(err).Msg("failed to set read deadline")
		fmt.Println(err)
	}

	err = ws.service.SetNewConnection(clientInfo, connection.DataSender(&connection.WebSocketConnection{Conn: conn}))
	if err != nil {
		ws.logger.Error().Err(err).Msg("failed to set new connection")
	}

	pingTicker := time.NewTicker(pingInterval)
	defer pingTicker.Stop()

	ctx, cancel := context.WithCancel(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-pingTicker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				ws.logger.Error().Err(err).Msg("ping failed")
				return
			}
		default:
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					ws.logger.Error().Msg("client disconnected normally")
					return
				}
				ws.logger.Error().Err(err).Msg("readMessage failed")
				return
			}

			if messageType == websocket.PingMessage || messageType == websocket.PongMessage {
				continue
			}

			if len(message) == 0 {
				continue
			}

			if ws.rateLimit.IsLimited(token) {
				ws.logger.Warn().Msg("rate limit exceeded")
				continue
			}

			ds := connection.DataSender(&connection.WebSocketConnection{Conn: conn, CtxClose: cancel})
			users, newMessage, err := ws.service.GetUsersAndMessage(ds, clientInfo, message)
			if err != nil {
				if err := conn.WriteMessage(websocket.TextMessage, []byte(err.Error())); err != nil {
					ws.logger.Error().Err(err).Msg("failed to send error message to client")
				}
				continue
			}

			for _, user := range users {
				if err := user.Connection.Write(newMessage); err != nil {
					ws.logger.Error().Err(err).Msg("failed to send message to user")
					if err := ws.service.RemoveUser(clientInfo, user.ID); err != nil {
						ws.logger.Error().Err(err).Msg("failed to remove user")
					}
				}
			}
		}
	}
}

func (ws *WebSocket) Run(addr string) error {
	router := mux.NewRouter()
	router.HandleFunc("/api/ws/connect", ws.connect)

	ws.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return ws.server.ListenAndServe()
}

func NewWebSocket(service service.Service, rateLimit utils.RateLimit, logger zerolog.Logger) *WebSocket {
	return &WebSocket{
		service:   service,
		rateLimit: rateLimit,
		mtx:       sync.RWMutex{},
		logger:    logger,
	}
}
