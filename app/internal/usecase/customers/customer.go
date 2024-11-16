package customers

import (
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/logging"
)

type CustomerUseCase struct {
	repository interfaces.CustomerRepository
	logger     logging.Logger
}

func NewCustomerUseCase(repository interfaces.CustomerRepository, logger logging.Logger) *CustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
		logger:     logger,
	}
}
