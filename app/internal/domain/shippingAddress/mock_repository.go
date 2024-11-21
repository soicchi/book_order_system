package shippingAddress

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockShippingAddressRepository struct {
	mock.Mock
}

func NewMockShippingAddressRepository() *MockShippingAddressRepository {
	return &MockShippingAddressRepository{}
}

func (m *MockShippingAddressRepository) Create(ctx echo.Context, shippingAddress *ShippingAddress) error {
	args := m.Called(ctx, shippingAddress)
	return args.Error(0)
}

func (m *MockShippingAddressRepository) FetchByID(ctx echo.Context, id string) (*ShippingAddress, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*ShippingAddress), args.Error(1)
}
