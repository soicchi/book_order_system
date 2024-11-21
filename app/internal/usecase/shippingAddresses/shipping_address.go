package shippingAddresses

import (
	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/domain/shippingAddress"
	"github.com/soicchi/book_order_system/internal/logging"
)

type ShippingAddressUseCase struct {
	shippingAddressRepo shippingAddress.Repository
	customerRepo        customer.Repository
	logger              logging.Logger
}

func NewShippingAddressUseCase(
	shippingAddressRepo shippingAddress.Repository,
	customerRepo customer.Repository,
	logger logging.Logger,
) *ShippingAddressUseCase {
	return &ShippingAddressUseCase{
		shippingAddressRepo: shippingAddressRepo,
		customerRepo:        customerRepo,
		logger:              logger,
	}
}
