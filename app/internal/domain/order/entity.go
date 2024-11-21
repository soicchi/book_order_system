package order

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Order struct {
	id        uuid.UUID
	createdAt *time.Time
	updatedAt *time.Time
}

func New() (*Order, error) {
	orderUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to generate order uuid: %w", err),
			errors.InternalServerError,
		)
	}

	return new(orderUUID, nil, nil), nil
}

func new(id uuid.UUID, createdAt, updatedAt *time.Time) *Order {
	return &Order{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) CreatedAt() *time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() *time.Time {
	return o.updatedAt
}
