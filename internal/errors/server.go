package errors

import "errors"

var (
	ErrServerNotExists                     = errors.New("server not exists")
	ErrServerFullOfConnections             = errors.New("server full of connection")
	ErrServerCreatingRoomError             = errors.New("error creating room")
	ErrServerCreatingRoomAllServesOffError = errors.New("error creating room all serve off")
	ErrRoomIsExists                        = errors.New("room is exists")
	ErrRoomNoFound                         = errors.New("room not found")
)
