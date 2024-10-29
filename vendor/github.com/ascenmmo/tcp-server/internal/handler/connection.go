package handler

import (
	"context"
	"github.com/ascenmmo/tcp-server/internal/service"
	"github.com/ascenmmo/tcp-server/internal/utils"
	"github.com/ascenmmo/tcp-server/pkg/errors"
	"github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	"github.com/google/uuid"
)

type RestConnection struct {
	rateLimit utils.RateLimit
	server    service.Service
}

func (r *RestConnection) SetSendMessage(ctx context.Context, token string, message types.RequestSetMessage) (err error) {
	limited := r.rateLimit.IsLimited(token)
	if limited {
		return errors.ErrTooManyRequests
	}

	err = r.server.SetMessage(token, message)

	return err
}

func (r *RestConnection) GetMessage(ctx context.Context, token string) (messages types.ResponseGetMessage, err error) {
	limited := r.rateLimit.IsLimited(token)
	if limited {
		return messages, errors.ErrTooManyRequests
	}
	messages, err = r.server.GetMessages(token)
	return
}

func (r *RestConnection) RemoveUser(ctx context.Context, token string, userID uuid.UUID) (err error) {
	err = r.server.RemoveUser(userID, token)
	return
}

func NewRestConnection(rateLimit utils.RateLimit, server service.Service) *RestConnection {
	return &RestConnection{rateLimit: rateLimit, server: server}
}
