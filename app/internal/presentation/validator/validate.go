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
		return ers.New(
			fmt.Errorf("validation error: %w", err),
			ers.ValidationError,
		)
	}

	return nil
}
