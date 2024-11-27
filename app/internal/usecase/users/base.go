package users

import (
	"github.com/soicchi/book_order_system/internal/domain/user"
	"github.com/soicchi/book_order_system/internal/logging"
)

type UserUseCase struct {
	userRepository user.Repository
	logger         logging.Logger
}

func NewUseCase(userRepository user.Repository, logger logging.Logger) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
		logger:         logger,
	}
}
