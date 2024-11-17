package shippingAddresses

import (
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/logging"
)

type ShippingAddressUseCase struct {
	shippingAddressRepo interfaces.ShippingAddressRepository
	customerRepo        interfaces.CustomerRepository
	logger              logging.Logger
}

func NewShippingAddressUseCase(
	shippingAddressRepo interfaces.ShippingAddressRepository,
	customerRepo interfaces.CustomerRepository,
	logger logging.Logger,
) *ShippingAddressUseCase {
	return &ShippingAddressUseCase{
		shippingAddressRepo: shippingAddressRepo,
		customerRepo:        customerRepo,
		logger:              logger,
	}
}
