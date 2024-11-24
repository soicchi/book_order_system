package users

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (uu *UserUseCase) RetrieveUser(ctx echo.Context, id uuid.UUID) (*RetrieveOutput, error) {
	u, err := uu.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return NewRetrieveOutput(
		u.ID(),
		u.Username(),
		u.Email(),
		u.CreatedAt().Format("2006-01-02 15:04:05"),
		u.UpdatedAt().Format("2006-01-02 15:04:05"),
	), nil
}
