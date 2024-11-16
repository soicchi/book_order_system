package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorCode int

const (
	InvalidRequest ErrorCode = iota
	NotAuthorized
	NotFound
	Forbidden
	AlreadyExist
	InternalServerError
)

func (e ErrorCode) String() string {
	switch e {
	case InvalidRequest:
		return "invalid_request"
	case NotAuthorized:
		return "authorized_error"
	case NotFound:
		return "not_found"
	case Forbidden:
		return "forbidden"
	case AlreadyExist:
		return "already_exist"
	case InternalServerError:
		return "internal_error"
	default:
		return "internal_error"
	}
}

func (e ErrorCode) Status() int {
	switch e {
	case InvalidRequest:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case AlreadyExist:
		return http.StatusConflict
	case InternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (e ErrorCode) Message() string {
	switch e {
	case InvalidRequest:
		return "Invalid request parameters"
	case NotFound:
		return "Resource not found"
	case AlreadyExist:
		return "Resource already exists"
	default:
		return "An internal server error occurred"
	}
}

type Details map[string]interface{}

type CustomError struct {
	Err     error
	Message string
	Code    ErrorCode
	Details Details
}

type option func(*CustomError)

func WithDetails(details Details) option {
	return func(e *CustomError) {
		e.Details = details
	}
}

func WithNotFoundDetails(key string) option {
	return func(e *CustomError) {
		e.Details = Details{key: "not found"}
	}
}

func NewCustomError(err error, code ErrorCode, options ...option) *CustomError {
	customErr := &CustomError{
		Err:     err,
		Code:    code,
		Message: code.Message(),
		Details: nil,
	}

	for _, opt := range options {
		opt(customErr)
	}

	return customErr
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

// To return the wrapped error
func (c *CustomError) Unwrap() error {
	return c.Err
}

func (c *CustomError) ReturnJSON(ctx echo.Context) error {
	response := newErrorResponse(c.Code.String(), c.Message, c.Details)

	return ctx.JSON(c.Code.Status(), response)
}

type ErrorResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Details Details `json:"details,omitempty"`
}

func newErrorResponse(code, message string, details Details) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}
