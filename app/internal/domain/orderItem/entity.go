package orderItem

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type OrderItem struct {
	id       uuid.UUID
	quantity int
}

func New(quantity int) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, errors.New(
			fmt.Errorf("quantity must be greater than 0. got: %d", quantity),
			errors.InvalidRequest,
		)
	}

	orderItemUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New(
			fmt.Errorf("failed to generate order item UUID: %w", err),
			errors.InternalServerError,
		)
	}

	return new(orderItemUUID, quantity), nil
}

func new(id uuid.UUID, quantity int) *OrderItem {
	return &OrderItem{
		id:       id,
		quantity: quantity,
	}
}
