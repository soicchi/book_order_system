package customer

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	id  uuid.UUID
	name string
	email string
	prefecture string
	address string
	password string
}

func NewCustomer(name, email, prefecture, address, password string) (*Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Customer{
		id: uuid.New(),
		name: name,
		email: email,
		prefecture: prefecture,
		address: address,
		password: string(hashedPassword),
	}, nil
}
