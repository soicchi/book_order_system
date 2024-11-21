package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/shippingAddress"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderRepo           order.Repository
	customerRepo        customer.Repository
	shippingAddressRepo shippingAddress.Repository
	logger              logging.Logger
}

func NewOrderUseCase(
	orderRepo order.Repository,
	customerRepo customer.Repository,
	shippingAddressRepo shippingAddress.Repository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:           orderRepo,
		customerRepo:        customerRepo,
		shippingAddressRepo: shippingAddressRepo,
		logger:              logger,
	}
}
