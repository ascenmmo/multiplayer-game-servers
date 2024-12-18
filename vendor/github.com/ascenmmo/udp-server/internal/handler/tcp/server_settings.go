package tcp

import (
	"context"
	"github.com/ascenmmo/udp-server/internal/service"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/ascenmmo/udp-server/pkg/api/types"
	"github.com/ascenmmo/udp-server/pkg/errors"
	"github.com/google/uuid"
)

type ServerSettings struct {
	rateLimit utils.RateLimit
	server    service.Service
}

func (r *ServerSettings) GetConnectionsNum(ctx context.Context, token string) (countConn int, exists bool, err error) {
	limited := r.rateLimit.IsLimited(token)
	if limited {
		return countConn, exists, errors.ErrTooManyRequests
	}
	countConn, exists = r.server.GetConnectionsNum()
	return
}

func (r *ServerSettings) HealthCheck(ctx context.Context, token string) (exists bool, err error) {
	limited := r.rateLimit.IsLimited(token)
	if limited {
		return exists, errors.ErrTooManyRequests
	}
	return true, nil
}

func (r *ServerSettings) GetServerSettings(ctx context.Context, token string) (settings types.Settings, err error) {
	limited := r.rateLimit.IsLimited(token)
	if limited {
		return settings, errors.ErrTooManyRequests
	}

	return types.NewSettings(), nil
}

func (r *ServerSettings) CreateRoom(ctx context.Context, token string, createRoom types.CreateRoomRequest) (err error) {
	limited := r.rateLimit.IsLimited(token)
	if limited {
		return errors.ErrTooManyRequests
	}
	err = r.server.CreateRoom(token)
	return
}

func (r *ServerSettings) SetNotifyServer(ctx context.Context, token string, id uuid.UUID, url string) (err error) {
	return
}

func NewServerSettings(rateLimit utils.RateLimit, server service.Service) *ServerSettings {
	return &ServerSettings{rateLimit: rateLimit, server: server}
}
