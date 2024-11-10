package interfaces

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (r *CustomerRepositoryMock) Create(ctx echo.Context, customer *entity.Customer) error {
	args := r.Called(customer)
	return args.Error(0)
}
