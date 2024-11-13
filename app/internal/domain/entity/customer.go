package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Customer struct {
	id        uuid.UUID
	name      string
	email     string
	password  string
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

	if len(plainPassword) < 8 {
		return nil, errors.NewCustomError(
			fmt.Errorf("password must be at least 8 characters"),
			errors.InvalidRequest,
		)
	}

	// convert plain password to sh256 hash
	sha256Hash := sha256.Sum256([]byte(plainPassword))
	hashedPassword := hex.EncodeToString(sha256Hash[:])

	return newCustomer(customerUUID, name, email, hashedPassword, nil, nil), nil
}

func ReconstructCustomer(
	id uuid.UUID,
	name, email string,
	password string,
	createdAt *time.Time,
	updatedAt *time.Time,
) *Customer {
	return newCustomer(id, name, email, password, createdAt, updatedAt)
}

func newCustomer(id uuid.UUID, name, email, password string, createdAt, updatedAt *time.Time) *Customer {
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

func (c *Customer) Password() string {
	return c.password
}

func (c *Customer) CreatedAt() *time.Time {
	return c.createdAt
}

func (c *Customer) UpdatedAt() *time.Time {
	return c.updatedAt
}
