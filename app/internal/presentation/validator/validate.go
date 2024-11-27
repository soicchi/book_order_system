package validator

import (
	"fmt"

	ers "github.com/soicchi/book_order_system/internal/errors"

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

// TODO: 特定のFieldとIssueを返すようにする
// 一旦バリデーションエラーだけ返す
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// https://github.com/go-playground/validator/blob/6c3307e6c64040ebc0efffe9366927c68146ffba/errors.go#L80
		for _, fieldError := range err.(validator.ValidationErrors) {
			return ers.New(
				fmt.Errorf("validation error: %s. got: %s", fieldError.Field(), fieldError.Tag()),
				ers.ValidationError,
				ers.WithField(fieldError.Field()),
				ers.WithIssue(TagToErrorIssue(fieldError.Tag())),
			)
		}
	}

	return nil
}

func TagToErrorIssue(tag string) ers.ErrorIssue {
	switch tag {
	case "required":
		return ers.Required
	case "email":
		return ers.Invalid
	case "uuid":
		return ers.Invalid
	default:
		return ers.Unknown
	}
}
