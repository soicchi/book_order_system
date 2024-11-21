package customers

import (
	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/logging"
)

type CustomerUseCase struct {
	repository customer.CustomerRepository
	logger     logging.Logger
}

func NewCustomerUseCase(repository customer.Repository, logger logging.Logger) *CustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
		logger:     logger,
	}
}
