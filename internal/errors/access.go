package errors

import "errors"

var (
	ErrAccessDenied                = errors.New("access denied")
	ErrAccessDeniedDeleteCreatorID = errors.New("access denied cannot delete creatorID")
)
