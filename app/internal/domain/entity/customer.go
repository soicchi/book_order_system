package entity

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Customer struct {
	id        uuid.UUID
	name      string
	email     string
	password  values.Password
	createdAt *time.Time
	updatedAt *time.Time
}

func NewCustomer(name, email, plainPassword string) (*Customer, error) {
	customerUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to generate customer UUID: %w", err),
			errors.InternalServerError,
		)
	}

	hashedPassword, err := values.NewPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	return newCustomer(customerUUID, name, email, hashedPassword, nil, nil), nil
}

func ReconstructCustomer(
	id uuid.UUID,
	name string,
	email string,
	password values.Password,
	createdAt time.Time,
	updatedAt time.Time,
) *Customer {
	return newCustomer(id, name, email, password, &createdAt, &updatedAt)
}

func newCustomer(
	id uuid.UUID,
	name string,
	email string,
	password values.Password,
	createdAt *time.Time,
	updatedAt *time.Time,
) *Customer {
	return &Customer{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (c *Customer) ID() uuid.UUID {
	return c.id
}

func (c *Customer) Name() string {
	return c.name
}

func (c *Customer) Email() string {
	return c.email
}

func (c *Customer) Password() values.Password {
	return c.password
}

func (c *Customer) CreatedAt() *time.Time {
	return c.createdAt
}

func (c *Customer) UpdatedAt() *time.Time {
	return c.updatedAt
}
