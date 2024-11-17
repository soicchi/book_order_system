package validator

import (
	"errors"
	"fmt"

	er "github.com/soicchi/book_order_system/internal/errors"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return er.NewCustomError(
			fmt.Errorf("validation error: %w", err),
			er.InvalidRequest,
			er.WithDetails(cv.generateDetail(err)),
		)
	}

	return nil
}

func (cv *CustomValidator) generateDetail(err error) er.Details {
	details := make(er.Details, len(err.(validator.ValidationErrors)))
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			details[fieldError.Field()] = fieldError.Tag()
		}
	}

	return details
}
