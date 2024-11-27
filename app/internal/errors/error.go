package errors

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type CustomError struct {
	Err   error
	Code  ErrorCode
	Field string
	Issue ErrorIssue
}

type option func(*CustomError)

func WithField(field string) option {
	return func(options *CustomError) {
		options.Field = field
	}
}

func WithIssue(issue ErrorIssue) option {
	return func(options *CustomError) {
		options.Issue = issue
	}
}

func New(err error, code ErrorCode, opts ...option) *CustomError {
	customErr := &CustomError{
		Err:   err,
		Code:  code,
		Field: "",
		Issue: NoIssue,
	}

	for _, opt := range opts {
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

func (c *CustomError) ErrorCode() string {
	if c.Field != "" && c.Issue != NoIssue {
		return fmt.Sprintf("%s.%s.%s", c.Code.String(), c.Field, c.Issue.String())
	}

	// If there is no issue, return the error code with the field.
	// ex) "NotFound.BookID"
	if c.Field != "" {
		return fmt.Sprintf("%s_%s", c.Code.String(), c.Field)
	}

	return c.Code.String()
}

func ReturnJSON(ctx echo.Context, err error) error {
	customErr, ok := err.(*CustomError)
	if !ok {
		customErr = New(err, UnexpectedError)
	}

	res := newErrorResponse(customErr.ErrorCode(), customErr.Code.Message())

	return ctx.JSON(customErr.Code.Status(), res)
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}
