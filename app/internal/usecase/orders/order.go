package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderRepo           interfaces.OrderRepository
	customerRepo        interfaces.CustomerRepository
	shippingAddressRepo interfaces.ShippingAddressRepository
	logger              logging.Logger
}

func NewOrderUseCase(
	orderRepo interfaces.OrderRepository,
	customerRepo interfaces.CustomerRepository,
	shippingAddressRepo interfaces.ShippingAddressRepository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:           orderRepo,
		customerRepo:        customerRepo,
		shippingAddressRepo: shippingAddressRepo,
		logger:              logger,
	}
}
