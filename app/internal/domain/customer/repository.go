package customer

import (
	"github.com/labstack/echo/v4"
)

type CustomerRepository interface {
	Create(ctx echo.Context, customer *Customer) error
	FetchByID(ctx echo.Context, id string) (*Customer, error)
}
