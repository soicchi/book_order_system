package orderdetail

import (
	"github.com/labstack/echo/v4"
)

type Repository interface {
	BulkCreate(ctx echo.Context, orderDetails []*OrderDetail) error
}
