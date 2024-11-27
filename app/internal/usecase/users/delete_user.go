package users

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (uu *UserUseCase) DeleteUser(ctx echo.Context, id uuid.UUID) error {
	u, err := uu.userRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.New(
			fmt.Errorf("user not found"),
			errors.NotFoundError,
			errors.WithField("User"),
		)
	}

	return uu.userRepository.Delete(ctx, u.ID())
}
