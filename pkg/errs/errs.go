package errs

import (
	"errors"
	"net/http"
)

type APIError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *APIError) Unwrap() error {
	return e.Err
}

func New(code int, message string, err error) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NotFound(msg string) *APIError {
	return New(http.StatusNotFound, msg, errors.New(msg))
}

func Internal(err error) *APIError {
	return New(http.StatusInternalServerError, "internal server error", err)
}

func BadRequest(msg string) *APIError {
	return New(http.StatusBadRequest, msg, errors.New(msg))
}

func Unauthorized(msg string) *APIError {
	return New(http.StatusUnauthorized, msg, errors.New(msg))
}
