package shippingAddress

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type ShippingAddress struct {
	id         uuid.UUID
	prefecture string
	city       string
	state      string
	createdAt  *time.Time
	updatedAt  *time.Time
}

func New(prefecture, city, state string) (*ShippingAddress, error) {
	shippingAddressUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to generate shipping address UUID: %w", err),
			errors.InternalServerError,
		)
	}

	return new(shippingAddressUUID, prefecture, city, state, nil, nil), nil
}

func Reconstruct(
	id uuid.UUID,
	prefecture string,
	city string,
	state string,
	createdAt *time.Time,
	updatedAt *time.Time,
) *ShippingAddress {
	return new(id, prefecture, city, state, createdAt, updatedAt)
}

func new(
	id uuid.UUID,
	prefecture string,
	city string,
	state string,
	createdAt *time.Time,
	updatedAt *time.Time,
) *ShippingAddress {
	return &ShippingAddress{
		id:         id,
		prefecture: prefecture,
		city:       city,
		state:      state,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}
}

func (sa *ShippingAddress) ID() uuid.UUID {
	return sa.id
}

func (sa *ShippingAddress) Prefecture() string {
	return sa.prefecture
}

func (sa *ShippingAddress) City() string {
	return sa.city
}

func (sa *ShippingAddress) State() string {
	return sa.state
}

func (sa *ShippingAddress) CreatedAt() *time.Time {
	return sa.createdAt
}

func (sa *ShippingAddress) UpdatedAt() *time.Time {
	return sa.updatedAt
}