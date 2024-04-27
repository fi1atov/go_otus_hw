package structs

import "errors"

var (
	ErrNotFound = errors.New("record not found")
	ErrInternal = errors.New("internal error")
)
