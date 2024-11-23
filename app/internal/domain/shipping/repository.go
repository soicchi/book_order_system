package shipping

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, shipping *Shipping, orderID uuid.UUID) error
	UpdateStatus(ctx echo.Context, shipping *Shipping) error
}
