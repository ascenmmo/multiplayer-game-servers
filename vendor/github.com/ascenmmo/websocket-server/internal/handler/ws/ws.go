package wsconnection

import (
	"context"
	"encoding/json"
	"fmt"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/ascenmmo/websocket-server/internal/connection"
	"github.com/ascenmmo/websocket-server/internal/service"
	"github.com/ascenmmo/websocket-server/internal/utils"
	"github.com/ascenmmo/websocket-server/pkg/restconnection/types"
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
	writeTimeout = 10 * time.Second
	pingInterval = 30 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  64 * 1024 * 100,
	WriteBufferSize: 64 * 1024 * 100,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocket struct {
	server          *http.Server
	mtx             sync.RWMutex
	service         service.Service
	rateLimit       utils.RateLimit
	rateLimitBadMsg utils.RateLimit
	logger          zerolog.Logger
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
		ws.logger.Error().Err(err).Msg("Failed to upgrade connection")
		return
	}

	ws.handleConnection(context.Background(), token, clientInfo, conn)
}

func (ws *WebSocket) handleConnection(ctx context.Context, token string, clientInfo tokentype.Info, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		err := ws.service.RemoveUser(token, clientInfo.UserID)
		if err != nil {
			ws.logger.Error().Err(err).Msg("Failed to remove user")
		}
	}()

	conn.SetReadLimit(bufferSize)
	err := conn.SetReadDeadline(time.Now().Add(readTimeout))
	if err != nil {
		ws.logger.Error().Err(err).Msg("Failed to set read deadline")
		fmt.Println(err)
	}

	err = ws.service.SetNewConnection(clientInfo, connection.DataSender(&connection.WebSocketConnection{Conn: conn}))
	if err != nil {
		ws.logger.Error().Err(err).Msg("Failed to set new connection")
		fmt.Println(err)
	}

	pingTicker := time.NewTicker(pingInterval)
	defer pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-pingTicker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				ws.logger.Error().Err(err).Msg("Ping failed")
				return
			}
		default:
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					ws.logger.Info().Msg("Client disconnected normally")
				} else {
					ws.logger.Error().Err(err).Msg("ReadMessage failed")
				}
				return
			}

			if messageType == websocket.PingMessage || messageType == websocket.PongMessage {
				continue
			}

			if len(message) == 0 {
				continue
			}

			if ws.rateLimit.IsLimited(token) {
				ws.logger.Warn().Msg("Rate limit exceeded")
				continue
			}

			var request types.Request
			if err := json.Unmarshal(message, &request); err != nil {
				if ws.rateLimitBadMsg.IsLimited(token) {
					ws.logger.Warn().Msg("Rate limit exceeded for bad message formats")
					return
				}
				continue
			}
			ds := connection.DataSender(&connection.WebSocketConnection{Conn: conn})

			if request.Server != nil {
				clientInfo, err = ws.service.ParseToken(request.Token)
				if err != nil {
					ws.logger.Error().Err(err).Msg("Failed to parse token")
					return
				}
			}

			if request.Token != "" {
				request.Token = token
			}

			users, msg, err := ws.service.GetUsersAndMessage(ds, clientInfo, request)
			if err != nil {
				if err := conn.WriteMessage(websocket.TextMessage, []byte(err.Error())); err != nil {
					ws.logger.Error().Err(err).Msg("Failed to send error message to client")
				}
				continue
			}

			for _, user := range users {
				if err := user.Connection.Write(msg); err != nil {
					ws.logger.Error().Err(err).Msg("Failed to send message to user")
					if err := ws.service.RemoveUser(token, user.ID); err != nil {
						ws.logger.Error().Err(err).Msg("Failed to remove user")
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

func NewWebSocket(service service.Service, rateLimit, rateLimitBadMsg utils.RateLimit, logger zerolog.Logger) *WebSocket {
	return &WebSocket{
		service:         service,
		rateLimit:       rateLimit,
		rateLimitBadMsg: rateLimitBadMsg,
		mtx:             sync.RWMutex{},
		logger:          logger,
	}
}
