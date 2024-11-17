package shippingAddresses

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

func (u *ShippingAddressUseCase) CreateShippingAddress(ctx echo.Context, shippingAddress *dto.CreateShippingAddressInput) error {
	shippingAddressEntity, err := entity.NewShippingAddress(
		shippingAddress.Prefecture,
		shippingAddress.City,
		shippingAddress.State,
		shippingAddress.CustomerID,
	)
	if err != nil {
		return err
	}

	// Check if customer exists
	customer, err := u.customerRepo.FetchByID(ctx, shippingAddress.CustomerID)
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

	return u.shippingAddressRepo.Create(ctx, shippingAddressEntity)
}
