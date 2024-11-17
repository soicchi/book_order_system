package shippingAddresses

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"
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
	if _, err := u.customerRepo.FetchByID(ctx, shippingAddress.CustomerID); err != nil {
		return err
	}

	if err := u.shippingAddressRepo.Create(ctx, shippingAddressEntity); err != nil {
		return err
	}

	return nil
}
