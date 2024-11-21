package customer

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (r *MockRepository) Create(ctx echo.Context, customer *Customer) error {
	args := r.Called(customer)
	return args.Error(0)
}

func (r *MockRepository) FetchByID(ctx echo.Context, id string) (*Customer, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Customer), args.Error(1)
}
