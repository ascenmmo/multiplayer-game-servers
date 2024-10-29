package errors

import "errors"

var (
	ErrUserNotFound              = errors.New("user not found")
	ErrRoomNotFound              = errors.New("room not found")
	ErrRoomIsExists              = errors.New("room is exists")
	ErrRoomBadValue              = errors.New("room bad value")
	ErrTooManyRequests           = errors.New("too many connection")
	ErrNotifyServerNotFound      = errors.New("err notify server not found")
	ErrNotifyServerNotValid      = errors.New("err notify server not valid")
	ErrGameConfigMarshalUserData = errors.New("err game config marshal user data")
	ErrGameResultsNotFound       = errors.New("game results not found")
)
