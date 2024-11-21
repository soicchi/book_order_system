package customer

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func NewMockCustomerRepository() *MockCustomerRepository {
	return &MockCustomerRepository{}
}

func (r *MockCustomerRepository) Create(ctx echo.Context, customer *Customer) error {
	args := r.Called(customer)
	return args.Error(0)
}

func (r *MockCustomerRepository) FetchByID(ctx echo.Context, id string) (*Customer, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Customer), args.Error(1)
}
