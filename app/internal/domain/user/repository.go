package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	Create(ctx echo.Context, user *User) error
	FetchByID(ctx echo.Context, userID uuid.UUID) (*User, error)
	FetchAll(ctx echo.Context) ([]*User, error)
	Update(ctx echo.Context, user *User) error
}
