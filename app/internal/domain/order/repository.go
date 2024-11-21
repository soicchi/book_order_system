package order

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OrderRepository interface {
	Create(ctx echo.Context, order *Order, customerID, shippingAddressID uuid.UUID) error
}
