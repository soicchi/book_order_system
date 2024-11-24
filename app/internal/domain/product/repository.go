package product

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, product *Product) error
	FetchAllByIDs(ctx echo.Context, ids []uuid.UUID) (Products, error)
}
