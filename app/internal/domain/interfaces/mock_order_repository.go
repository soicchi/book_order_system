package interfaces

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{}
}

func (m *MockOrderRepository) Create(ctx echo.Context, order *entity.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}
