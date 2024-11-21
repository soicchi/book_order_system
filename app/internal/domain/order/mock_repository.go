package order

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{}
}

func (m *MockOrderRepository) Create(ctx echo.Context, order *Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}
