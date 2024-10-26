package errors

import "errors"

var (
	ErrServerNotExists                     = errors.New("server not exists")
	ErrServerCreatingRoomError             = errors.New("error creating room")
	ErrServerCreatingRoomAllServesOffError = errors.New("error creating room all serve off")
	ErrRoomIsExists                        = errors.New("room is exists")
)
