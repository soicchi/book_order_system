package customers

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

func (uc *CustomerUseCase) FetchCustomer(ctx echo.Context, id string) (*dto.CustomerOutput, error) {
	customer, err := uc.repository.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.NewCustomError(
			fmt.Errorf("customer not found"),
			errors.NotFound,
			errors.WithNotFoundDetails("customer_id"),
		)
	}

	return dto.NewCustomerOutput(
		customer.ID().String(),
		customer.Name(),
		customer.Email(),
		customer.CreatedAt(),
		customer.UpdatedAt(),
	), nil
}
