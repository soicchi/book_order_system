package users

import (
	"github.com/labstack/echo/v4"
)

func (uu *UserUseCase) ListUser(ctx echo.Context) ([]*ListOutput, error) {
	users, err := uu.userRepository.FindAll(ctx)
	if err != nil {
		return []*ListOutput{}, err
	}

	output := make([]*ListOutput, 0, len(users))

	for _, user := range users {
		output = append(output, NewListOutput(
			user.ID(),
			user.Username(),
			user.Email(),
			user.CreatedAt().Format("2006-01-02 15:04:05"),
			user.UpdatedAt().Format("2006-01-02 15:04:05"),
		))
	}

	return output, nil
}
