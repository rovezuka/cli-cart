package storage_errors

import "errors"

var (
	ErrHasNoCart       = errors.New("has no cart")
	ErrInvalidSKUCount = errors.New("invalid sku count")
)
