package users

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/labstack/echo/v4"
)

func (uu *UserUseCase) UpdateUser(ctx echo.Context, dto *UpdateInput) error {
	u, err := uu.userRepository.FindByID(ctx, dto.ID)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.New(
			fmt.Errorf("user not found"),
			errors.NotFound,
		)
	}

	if err := u.Update(dto.Username, dto.Email, dto.Password); err != nil {
		return err
	}

	return uu.userRepository.Update(ctx, u)
}
