package orderItem

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, orderItem *OrderItem, orderID, productID uuid.UUID) error
}
