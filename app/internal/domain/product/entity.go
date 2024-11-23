package product

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type Product struct {
	id        uuid.UUID
	name      string
	price     float64
	createdAt *time.Time
	updatedAt *time.Time
}

func New(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New(
			fmt.Errorf("name must not be empty"),
			errors.InvalidRequest,
		)
	}

	if price <= 0 {
		return nil, errors.New(
			fmt.Errorf("price must be greater than 0: %f", price),
			errors.InvalidRequest,
		)
	}

	productUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New(
			fmt.Errorf("failed to generate product UUID: %w", err),
			errors.InternalServerError,
		)
	}

	now := time.Now()

	return new(productUUID, name, price, &now, &now), nil
}

func new(id uuid.UUID, name string, price float64, createdAt *time.Time, updatedAt *time.Time) *Product {
	return &Product{
		id:        id,
		name:      name,
		price:     price,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (p *Product) ID() uuid.UUID {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) CreatedAt() *time.Time {
	return p.createdAt
}

func (p *Product) UpdatedAt() *time.Time {
	return p.updatedAt
}
