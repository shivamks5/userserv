package errs

import "errors"

var (
	ErrNotFound     = errors.New("user not found")
	ErrInvalidField = errors.New("invalid data field")
	ErrBadRequest   = errors.New("invalid request")
)
