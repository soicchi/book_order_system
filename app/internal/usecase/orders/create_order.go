package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/labstack/echo/v4"
)

func (u *OrderUseCase) CreateOrder(ctx echo.Context, dto *dto.CreateOrderInput) error {
	order, err := order.New()
	if err != nil {
		return err
	}

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

	// Check if shipping address exists
	shippingAddress, err := u.shippingAddressRepo.FetchByID(ctx, dto.ShippingAddressID)
	if err != nil {
		return err
	}

	if shippingAddress == nil {
		return errors.NewCustomError(
			fmt.Errorf("shipping address not found"),
			errors.NotFound,
			errors.WithNotFoundDetails("shipping_address_id"),
		)
	}

	return u.orderRepo.Create(ctx, order, dto.CustomerID, dto.ShippingAddressID)
}