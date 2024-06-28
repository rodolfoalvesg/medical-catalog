package vos

import (
	"back-platform/app/gateway/http/rest/responses"
	"errors"
)

var (
	ErrPasswordMismatch = NewError("password mismatch")
	ErrUserAndPassEmpty = NewError("username and password are required")
	ErrUserNotFound     = NewError("user not found")
)

// NewError creates a new error with the given message.
func NewError(msg string) error {
	return &Error{msg}
}

// Error is a custom error type.
type Error struct {
	msg string
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.msg
}

func Errors(err error) responses.Response {
	switch {
	case errors.Is(err, ErrPasswordMismatch):
		return responses.Unauthorized(err)
	case errors.Is(err, ErrUserNotFound):
		return responses.NotFound(err)
	default:
		return responses.InternalServerError(err)
	}
}
