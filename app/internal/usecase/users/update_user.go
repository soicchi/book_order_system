package users

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/user"
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

	// Reconstruct関数を利用して、ユーザー情報を更新する
	// 既存のEntityを更新する場合は、Reconstruct関数を利用する
	updatedUser := user.Reconstruct(
		dto.ID,
		dto.Username,
		dto.Email,
		dto.Password,
		u.CreatedAt(),
		time.Now(),
	)

	return uu.userRepository.Update(ctx, updatedUser)
}
