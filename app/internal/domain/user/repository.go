package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(ctx echo.Context, user *User) error
	FindByID(ctx echo.Context, id uuid.UUID) (*User, error)
	FindAll(ctx echo.Context) ([]*User, error)
	Update(ctx echo.Context, user *User) error
	Delete(ctx echo.Context, id uuid.UUID) error
}
