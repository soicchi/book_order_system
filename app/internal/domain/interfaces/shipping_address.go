package interfaces

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"

	"github.com/labstack/echo/v4"
)

type ShippingAddressRepository interface {
	Create(ctx echo.Context, shippingAddress *entity.ShippingAddress) error
}
