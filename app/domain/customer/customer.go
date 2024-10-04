package customer

import (
	"fmt"

	"github.com/soicchi/book_order_system/domain/values"

	"github.com/google/uuid"
)

type Customer struct {
	id         uuid.UUID
	name       string
	email      values.Email
	prefecture values.Prefecture
	address    string
	password   values.Password
}

func NewCustomer(name, email, prefecture, address, password string) (*Customer, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	emailObj := values.Email(email)
	if err := emailObj.Validate(); err != nil {
		return nil, fmt.Errorf("invalid email: %w", err)
	}

	prefectureObj := values.Prefecture(prefecture)
	if err := prefectureObj.Validate(); err != nil {
		return nil, fmt.Errorf("invalid prefecture: %w", err)
	}

	if address == "" {
		return nil, fmt.Errorf("address is required")
	}

	passwordObj := values.Password(password)
	if err := passwordObj.Validate(); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	hashedPassword, err := passwordObj.ToHashed()
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return &Customer{
		id:         uuid.New(),
		name:       name,
		email:      emailObj,
		prefecture: prefectureObj,
		address:    address,
		password:   hashedPassword,
	}, nil
}
