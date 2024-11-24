package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/txmanager"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderService order.OrderService
	txManager    txmanager.Repository
	logger       logging.Logger
}

func NewOrderUseCase(
	orderService order.OrderService,
	txManager txmanager.Repository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderService: orderService,
		txManager:    txManager,
		logger:       logger,
	}
}
