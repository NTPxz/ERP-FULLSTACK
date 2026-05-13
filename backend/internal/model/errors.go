package model

import "errors"

var (
	ErrNotFound   = errors.New("resource not found")
	ErrBadRequest = errors.New("invalid request")
	ErrConflict   = errors.New("resource already exists")
	ErrForbidden  = errors.New("access denied")
)
