package errors

import "errors"

var (
	ErrGameConfigSameResultName = errors.New("error game config same result name")
	ErrGameConfigSameColumnName = errors.New("error game config same column name")
	ErrGameConfigNotFound       = errors.New("game config not found")
)
