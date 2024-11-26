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

func (m *MockRepository) Create(ctx echo.Context, order *Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockRepository) FindByID(ctx echo.Context, orderID uuid.UUID) (*Order, error) {
	args := m.Called(ctx, orderID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Order), args.Error(1)
}

func (m *MockRepository) UpdateStatus(ctx echo.Context, order *Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}
