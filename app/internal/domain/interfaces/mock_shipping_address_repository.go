package interfaces

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockShippingAddressRepository struct {
	mock.Mock
}

func NewMockShippingAddressRepository() *MockShippingAddressRepository {
	return &MockShippingAddressRepository{}
}

func (m *MockShippingAddressRepository) Create(ctx echo.Context, shippingAddress *entity.ShippingAddress) error {
	args := m.Called(ctx, shippingAddress)
	return args.Error(0)
}

func (m *MockShippingAddressRepository) FetchByID(ctx echo.Context, id string) (*entity.ShippingAddress, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.ShippingAddress), args.Error(1)
}
