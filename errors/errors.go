package errors

import (
	"net/http"
)

type WrappedError interface {
	error
	StatusCode() int
	ErrCode() string
}

type ErrorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// InvalidMethod represent as the error for invalid method
type InvalidMethod string

func (e InvalidMethod) Error() string {
	return string(e)
}

func (e InvalidMethod) StatusCode() int {
	return http.StatusMethodNotAllowed
}

func (e InvalidMethod) ErrCode() string {
	return "METHOD_NOT_ALLOWED"
}

// InvalidInput represent as the error for invalid input
type InvalidInput string

func (e InvalidInput) Error() string {
	return string(e)
}

func (e InvalidInput) StatusCode() int {
	return http.StatusBadRequest
}

func (e InvalidInput) ErrCode() string {
	return "INVALID_INPUT"
}

// Registered Error
var (
	ErrInvalidMethod            = InvalidMethod("Please use GET and POST only for this endpoint")
	ErrInvalidInput             = InvalidInput("Please specify input field")
	ErrInvalidOperation         = InvalidInput("Operation is not allowed")
	ErrInvalidInputToBeOperated = InvalidInput("Input can't be operated")
)
