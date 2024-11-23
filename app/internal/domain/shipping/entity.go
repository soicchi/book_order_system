package shipping

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Shipping struct {
	id        uuid.UUID
	address   string
	method    values.ShippingMethod
	status    values.ShippingStatus
	createdAt *time.Time
	updatedAt *time.Time
}

func New(address string, method string) (*Shipping, error) {
	shippingUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New(
			fmt.Errorf("failed to generate uuid: %w", err),
			errors.InternalServerError,
		)
	}

	shippingMethod, err := values.NewShippingMethod("pending")
	if err != nil {
		return nil, fmt.Errorf("failed to create shipping method: %w", err)
	}
	now := time.Now()

	return new(shippingUUID, address, shippingMethod, values.Waiting, &now, &now), nil
}

func new(id uuid.UUID,
	address string,
	method values.ShippingMethod,
	status values.ShippingStatus,
	createdAt *time.Time,
	updatedAt *time.Time,
) *Shipping {
	return &Shipping{
		id:        id,
		address:   address,
		method:    method,
		status:    status,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (s *Shipping) ID() uuid.UUID {
	return s.id
}

func (s *Shipping) Address() string {
	return s.address
}

func (s *Shipping) Method() values.ShippingMethod {
	return s.method
}

func (s *Shipping) Status() values.ShippingStatus {
	return s.status
}

func (s *Shipping) CreatedAt() *time.Time {
	return s.createdAt
}

func (s *Shipping) UpdatedAt() *time.Time {
	return s.updatedAt
}
