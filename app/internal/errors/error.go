package errors

import (
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

func ReturnJSON(ctx echo.Context, err error) error {
	customErr, ok := err.(*CustomError)
	if !ok {
		customErr = New(err, UnexpectedError)
	}

	res := newErrorResponse(
		customErr.Code.String(),
		customErr.Field,
		customErr.Issue.String(),
		customErr.Code.Message(),
	)

	return ctx.JSON(customErr.Code.Status(), res)
}

type DetailResponse struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

func newDetailResponse(field, issue string) DetailResponse {
	return DetailResponse{
		Field: field,
		Issue: issue,
	}
}

type ErrorResponse struct {
	Code    string         `json:"code"`
	Detail  DetailResponse `json:"detail"`
	Message string         `json:"message"`
}

func newErrorResponse(code, field string, issue string, message string) ErrorResponse {
	detail := newDetailResponse(field, issue)

	return ErrorResponse{
		Code:    code,
		Detail:  detail,
		Message: message,
	}
}
