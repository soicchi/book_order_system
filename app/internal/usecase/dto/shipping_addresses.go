package dto

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type CreateShippingAddressInput struct {
	CustomerID uuid.UUID
	Prefecture string
	City       string
	State      string
}

func NewCreateShippingAddressInput(customerID, prefecture, city, state string) (*CreateShippingAddressInput, error) {
	customerUUID, err := uuid.Parse(customerID)
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to parse customer UUID: %w", err),
			errors.InvalidRequest,
		)
	}

	return &CreateShippingAddressInput{
		CustomerID: customerUUID,
		Prefecture: prefecture,
		City:       city,
		State:      state,
	}, nil
}
