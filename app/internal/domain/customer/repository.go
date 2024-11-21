package customer

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, customer *Customer) error
	FetchByID(ctx echo.Context, id uuid.UUID) (*Customer, error)
}
