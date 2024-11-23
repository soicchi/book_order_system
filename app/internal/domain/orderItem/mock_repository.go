package orderItem

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

func (r *MockRepository) Create(ctx echo.Context, orderItem *OrderItem, orderID, productID uuid.UUID) error {
	args := r.Called(ctx, orderItem, orderID, productID)
	return args.Error(0)
}
