package errors

import "errors"

var (
	ErrClientCreationError = errors.New("error creating client")
	ErrClientNotFound      = errors.New("error client not found")
)
