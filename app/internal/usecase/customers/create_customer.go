package customers

import (
	customerDomain "github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/dto"

	"github.com/labstack/echo/v4"
)

func (u *CustomerUseCase) CreateCustomer(ctx echo.Context, input *dto.CreateCustomerInput) error {
	// convert input to domain entity
	customer, err := customerDomain.New(input.Name, input.Email)
	if err != nil {
		return err
	}

	if err := u.customerRepo.Create(ctx, customer); err != nil {
		return err
	}

	return nil
}
