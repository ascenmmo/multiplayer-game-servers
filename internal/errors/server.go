package errors

import "errors"

var (
	ErrServerNotExists         = errors.New("server not exists")
	ErrServerCreatingRoomError = errors.New("error creating room")
)
