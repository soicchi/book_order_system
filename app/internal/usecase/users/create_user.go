package users

import (
	"github.com/soicchi/book_order_system/internal/domain/user"

	"github.com/labstack/echo/v4"
)

func (uu *UserUseCase) CreateUser(ctx echo.Context, dto *CreateInput) error {
	u, err := user.New(dto.Username, dto.Email, dto.Password)
	if err != nil {
		return err
	}

	return uu.userRepository.Create(ctx, u)
}
