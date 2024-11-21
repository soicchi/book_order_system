package customers

import (
	"time"

	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

func (uc *CustomerUseCase) CreateCustomer(ctx echo.Context, dto *dto.CreateCustomerInput) error {
	customer, err := entity.New(dto.Name, dto.Email, dto.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return uc.repository.Create(ctx, customer)
}
