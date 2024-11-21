package shippingAddresses

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/shippingAddress"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

func (u *ShippingAddressUseCase) CreateShippingAddress(ctx echo.Context, dto *dto.CreateShippingAddressInput) error {
	// Check if customer exists
	customer, err := u.customerRepo.FetchByID(ctx, dto.CustomerID)
	if err != nil {
		return err
	}
	if customer == nil {
		return errors.NewCustomError(
			fmt.Errorf("customer not found"),
			errors.NotFound,
			errors.WithNotFoundDetails("customer_id"),
		)
	}

	shippingAddressEntity, err := shippingAddress.New(
		dto.Prefecture,
		dto.City,
		dto.State,
	)
	if err != nil {
		return err
	}

	return u.shippingAddressRepo.Create(ctx, shippingAddressEntity, dto.CustomerID)
}
