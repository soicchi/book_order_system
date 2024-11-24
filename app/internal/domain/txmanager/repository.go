package txmanager

import (
	"github.com/labstack/echo/v4"
)

type Repository interface {
	WithTransaction(ctx echo.Context, fn func(ctx echo.Context) error) error
}
