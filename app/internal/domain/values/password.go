package values

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type Password string

func NewPassword(plainPassword string) (Password, error) {
	if len(plainPassword) < 8 {
		return "", errors.NewCustomError(
			fmt.Errorf("password must be at least 8 characters"),
			errors.InvalidRequest,
		)
	}

	// convert plain password to sh256 hash
	sha256Hash := sha256.Sum256([]byte(plainPassword))
	hashedPassword := hex.EncodeToString(sha256Hash[:])

	return Password(hashedPassword), nil
}

func (p Password) String() string {
	return string(p)
}
