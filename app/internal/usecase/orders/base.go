package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/domain/txmanager"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderService    order.OrderService
	orderDetailRepo orderdetail.Repository
	txManager       txmanager.Repository
	logger          logging.Logger
}

func NewOrderUseCase(
	orderService order.OrderService,
	orderDetailRepo orderdetail.Repository,
	txManager txmanager.Repository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderService:    orderService,
		orderDetailRepo: orderDetailRepo,
		txManager:       txManager,
		logger:          logger,
	}
}
