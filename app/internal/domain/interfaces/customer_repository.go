package interfaces

import (
	"github.com/soicchi/book_order_system/internal/domain/entity"

	"github.com/labstack/echo/v4"
)

type CustomerRepository interface {
	Create(ctx echo.Context, customer *entity.Customer) error
	FetchByID(ctx echo.Context, id string) (*entity.Customer, error)
}
