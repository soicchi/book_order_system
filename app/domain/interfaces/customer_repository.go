package interfaces

import (
	"github.com/soicchi/book_order_system/domain/entity"

	"github.com/labstack/echo/v4"
)

type CustomerRepository interface {
	Create(ctx echo.Context, customer *entity.Customer) error
}
