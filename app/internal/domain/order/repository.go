package order

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, order *Order, userID uuid.UUID) error
	FindByID(ctx echo.Context, orderID uuid.UUID) (*Order, error)
}
