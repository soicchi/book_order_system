package book

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, book *Book) error
	FindByID(ctx echo.Context, id uuid.UUID) (*Book, error)
	FindAll(ctx echo.Context) ([]*Book, error)
	Update(ctx echo.Context, book *Book) error
	Delete(ctx echo.Context, id uuid.UUID) error
}
