package orderItem

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	BulkCreate(ctx echo.Context, orderItem []*OrderItem, orderID uuid.UUID, productIDs []uuid.UUID) error
}
