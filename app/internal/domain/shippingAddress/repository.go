package shippingAddress

import (
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, shippingAddress *ShippingAddress, customerID string) error
	FetchByID(ctx echo.Context, id string) (*ShippingAddress, error)
}
