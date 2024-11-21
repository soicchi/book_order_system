package shippingAddress

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ShippingAddressRepository interface {
	Create(ctx echo.Context, shippingAddress *ShippingAddress, customerID uuid.UUID) error
	FetchByID(ctx echo.Context, id string) (*ShippingAddress, error)
}
