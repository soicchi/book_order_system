package orderdetail

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	BulkCreate(ctx echo.Context, orderDetails []*OrderDetail, orderID uuid.UUID) error
}
