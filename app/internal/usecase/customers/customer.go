package customers

import (
	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/logging"
)

type CustomerUseCase struct {
	customerRepo customer.Repository
	logger       logging.Logger
}

func NewCustomerUseCase(repository customer.Repository, logger logging.Logger) *CustomerUseCase {
	return &CustomerUseCase{
		customerRepo: repository,
		logger:       logger,
	}
}
