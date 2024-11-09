package values

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Password struct {
	value string
}

func NewPassword(plainPassword string) (*Password, error) {
	if len(plainPassword) < 8 {
		return nil, fmt.Errorf("password must be at least 8 characters")
	}

	// convert plain password to sh256 hash
	sha256Hash := sha256.Sum256([]byte(plainPassword))
	hashPassword := hex.EncodeToString(sha256Hash[:])

	return &Password{
		value: hashPassword,
	}, nil
}

func (p *Password) Value() string {
	return p.value
}
