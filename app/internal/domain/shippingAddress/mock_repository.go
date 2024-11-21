package shippingAddress

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

func (m *MockRepository) Create(ctx echo.Context, shippingAddress *ShippingAddress, customerID uuid.UUID) error {
	args := m.Called(ctx, shippingAddress, customerID)
	return args.Error(0)
}

func (m *MockRepository) FetchByID(ctx echo.Context, id uuid.UUID) (*ShippingAddress, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*ShippingAddress), args.Error(1)
}
