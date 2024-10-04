package values

import (
	"fmt"
	"net/mail"
)

type Email string

func (e Email) Validate() error {
	_, err := mail.ParseAddress(string(e))
	if err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	return nil
}
