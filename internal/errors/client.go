package errors

import "errors"

var (
	ErrClientCreationError = errors.New("error creating client")
	ErrNotFound            = errors.New("error not found")
	ErrWrongUserOrPassword = errors.New("error wrong user or password")
	ErrBadNewPassword      = errors.New("error bad new password")
	ErrGameSaves           = errors.New("error game saves found")
)
