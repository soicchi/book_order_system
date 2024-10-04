package values

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) Validate() error {
	if len(p) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	return nil
}

func (p Password) ToHashed() (Password, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(string(hashedPassword)), nil
}
