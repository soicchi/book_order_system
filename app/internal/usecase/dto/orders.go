package dto

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type CreateOrderInput struct {
	CustomerID        uuid.UUID
	ShippingAddressID uuid.UUID
}

func NewCreateOrderInput(customerID, shippingAddressID string) (*CreateOrderInput, error) {
	customerUUID, err := uuid.Parse(customerID)
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to parse customer UUID: %w", err),
			errors.InvalidRequest,
		)
	}

	shippingAddressUUID, err := uuid.Parse(shippingAddressID)
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to parse shipping address UUID: %w", err),
			errors.InvalidRequest,
		)
	}

	return &CreateOrderInput{
		CustomerID:        customerUUID,
		ShippingAddressID: shippingAddressUUID,
	}, nil
}
