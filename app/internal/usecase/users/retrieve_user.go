package users

import (
	"fmt"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (uu *UserUseCase) RetrieveUser(ctx echo.Context, id uuid.UUID) (*RetrieveOutput, error) {
	u, err := uu.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.New(
			fmt.Errorf("user not found"),
			errors.NotFoundError,
			errors.WithField("User"),
		)
	}

	return NewRetrieveOutput(
		u.ID(),
		u.Username(),
		u.Email(),
		u.CreatedAt().Format("2006-01-02 15:04:05"),
		u.UpdatedAt().Format("2006-01-02 15:04:05"),
	), nil
}
