package customer

import (
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, customer *Customer) error
}
