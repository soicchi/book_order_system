package errors

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
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

type ResourceType int

// define resource type here if other resources are added
const (
	Customer ResourceType = iota
)

func (r ResourceType) String() string {
	switch r {
	case Customer:
		return "customer"
	default:
		return "unknown"
	}
}

type CustomError struct {
	Err      error
	Code     ErrorCode
	Resource ResourceType
}

func NewCustomError(err error, code ErrorCode, resource ResourceType) *CustomError {
	return &CustomError{
		Err:      err,
		Code:     code,
		Resource: resource,
	}
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

func (c *CustomError) ReturnJSON(ctx echo.Context) error {
	var identifiers Identifiers
	var message string

	switch c.Code {
	case InvalidRequest:
		// handle invalid request error
		identifiers = c.handleInvalidRequest()
		message = "Invalid request parameters"
	// TODO: handler notfound error
	default:
		// handle other error
		identifiers = nil
		message = "An internal server error occurred"
	}

	if identifiers != nil {
		response := newErrorResponse(c.Code.String(), message, c.Resource.String(), WithIdentifiers(identifiers))
		return ctx.JSON(c.Code.Status(), response)
	}

	response := newErrorResponse(c.Code.String(), message, c.Resource.String())
	return ctx.JSON(c.Code.Status(), response)
}

func (c *CustomError) handleInvalidRequest() Identifiers {
	var identifiers Identifiers

	if errors.As(c.Err, &validator.ValidationErrors{}) {
		for _, fieldError := range c.Err.(validator.ValidationErrors) {
			identifiers[fieldError.Field()] = fieldError.Tag()
		}
	}

	return identifiers
}
