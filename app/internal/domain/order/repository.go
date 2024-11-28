package order

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, order *Order) error
	FindByID(ctx echo.Context, orderID uuid.UUID) (*Order, error)
	FindByIDWithOrderDetails(ctx echo.Context, orderID uuid.UUID) (*Order, error)
	UpdateStatus(ctx echo.Context, order *Order) error
}
