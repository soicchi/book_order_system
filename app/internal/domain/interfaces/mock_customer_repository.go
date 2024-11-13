package interfaces

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func NewMockCustomerRepository() *MockCustomerRepository {
	return &MockCustomerRepository{}
}

func (r *MockCustomerRepository) Create(ctx echo.Context, customer *entity.Customer) error {
	args := r.Called(customer)
	return args.Error(0)
}

func (r *MockCustomerRepository) FetchByID(ctx echo.Context, id string) (*entity.Customer, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.Customer), args.Error(1)
}
