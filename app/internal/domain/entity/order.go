package entity

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Order struct {
	id                uuid.UUID
	customerID        uuid.UUID
	shippingAddressID uuid.UUID
	createdAt         *time.Time
	updatedAt         *time.Time
}

func NewOrder(customerID, shippingAddressID string) (*Order, error) {
	orderUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to generate order uuid: %w", err),
			errors.InternalServerError,
		)
	}

	customerUUID, err := uuid.Parse(customerID)
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to parse customer uuid: %w", err),
			errors.InvalidRequest,
		)
	}

	shippingAddressUUID, err := uuid.Parse(shippingAddressID)
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to parse shipping address uuid: %w", err),
			errors.InvalidRequest,
		)
	}

	return newOrder(orderUUID, customerUUID, shippingAddressUUID, nil, nil), nil
}

func newOrder(id, customerID, shippingAddressID uuid.UUID, createdAt, updatedAt *time.Time) *Order {
	return &Order{
		id:                id,
		customerID:        customerID,
		shippingAddressID: shippingAddressID,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) CustomerID() uuid.UUID {
	return o.customerID
}

func (o *Order) ShippingAddressID() uuid.UUID {
	return o.shippingAddressID
}

func (o *Order) CreatedAt() *time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() *time.Time {
	return o.updatedAt
}
