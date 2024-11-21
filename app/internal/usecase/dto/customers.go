package dto

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

type CreateCustomerInput struct {
	Name     string
	Email    string
	Password string
}

type CustomerOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCustomerOutput(id, name, email string, createdAt, updatedAt time.Time) *CustomerOutput {
	return &CustomerOutput{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: updatedAt.Format("2006-01-02 15:04:05"),
	}
}

type FetchCustomerInput struct {
	CustomerID uuid.UUID
}

func NewFetchCustomerInput(id string) (*FetchCustomerInput, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("failed to parse customer UUID: %w", err),
			errors.InvalidRequest,
		)
	}

	return &FetchCustomerInput{
		CustomerID: uuid,
	}, nil
}
