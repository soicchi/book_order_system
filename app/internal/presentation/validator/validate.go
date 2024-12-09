package validator

import (
	"fmt"

	ers "event_system/internal/errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

func BindAndValidate(ctx echo.Context, i interface{}) error {
	// Echo の Bind 関数はJSON形式や型のチェックを行うものでユーザーからの入力を検証するものではないく、
	// どのフィールドがエラーなのかを判定することができないため
	// ValidationErrorとだけ返す。
	if err := ctx.Bind(i); err != nil {
		return ers.New(fmt.Errorf("failed to bind request: %w", err), ers.ValidationError)
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}
