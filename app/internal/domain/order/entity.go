package order

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Order struct {
	id          uuid.UUID
	totalAmount float64
	orderedAt   *time.Time
}

func New(totalAmount float64) (*Order, error) {
	orderUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New(
			fmt.Errorf("failed to generate uuid: %w", err),
			errors.InternalServerError,
		)
	}

	if totalAmount <= 0 {
		return nil, errors.New(
			fmt.Errorf("total amount must be greater than 0: %f", totalAmount),
			errors.InvalidRequest,
		)
	}

	now := time.Now()

	return new(orderUUID, totalAmount, &now), nil
}

func new(id uuid.UUID, totalAmount float64, orderedAt *time.Time) *Order {
	return &Order{
		id:          id,
		totalAmount: totalAmount,
		orderedAt:   orderedAt,
	}
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) TotalAmount() float64 {
	return o.totalAmount
}

func (o *Order) OrderedAt() *time.Time {
	return o.orderedAt
}
