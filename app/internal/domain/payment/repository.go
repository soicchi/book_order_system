package payment

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, payment *Payment, orderID uuid.UUID) error
	UpdateStatus(ctx echo.Context, payment *Payment) error
}
