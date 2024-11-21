package order

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) Create(ctx echo.Context, order *Order, customerID, shippingAddressID uuid.UUID) error {
	args := m.Called(ctx, order, customerID, shippingAddressID)
	return args.Error(0)
}
