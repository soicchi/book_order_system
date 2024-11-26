package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if len(value) == 0 {
		return Password{}, errors.New(
			fmt.Errorf("password must not be empty"),
			errors.InvalidRequest,
		)
	}

	// convert password to hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, errors.New(
			fmt.Errorf("failed to generate hash from password: %w", err),
			errors.InternalServerError,
		)
	}

	return Password{
		value: string(passwordHash),
	}, nil
}

func ReconstructPassword(value string) Password {
	return Password{
		value: value,
	}
}

func (p Password) Value() string {
	return p.value
}