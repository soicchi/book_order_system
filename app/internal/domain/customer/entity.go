package customer

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Customer struct {
	id        uuid.UUID
	name      string
	email     string
	createdAt *time.Time
	updatedAt *time.Time
}

func New(name, email string) (*Customer, error) {
	customerUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New(
			fmt.Errorf("failed to generate uuid: %w", err),
			errors.InternalServerError,
		)
	}

	now := time.Now()

	return new(customerUUID, name, email, &now, &now), nil
}

func new(id uuid.UUID, name, email string, createdAt, updatedAt *time.Time) *Customer {
	return &Customer{
		id:        id,
		name:      name,
		email:     email,
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

func (c *Customer) CreatedAt() *time.Time {
	return c.createdAt
}

func (c *Customer) UpdatedAt() *time.Time {
	return c.updatedAt
}
