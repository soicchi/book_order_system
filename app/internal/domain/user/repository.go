package user

import (
	"event_system/internal/domain/role"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	Create(ctx echo.Context, user *User) error
	FetchByID(ctx echo.Context, userID uuid.UUID) (*User, error)
	FetchByRole(ctx echo.Context, role *role.Role) ([]*User, error)
	Update(ctx echo.Context, user *User) error
}
