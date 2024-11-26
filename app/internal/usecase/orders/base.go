package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/book"
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/domain/txmanager"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderRepository       order.Repository
	bookService           book.Service
	orderDetailRepository orderdetail.Repository
	txManager             txmanager.Repository
	logger                logging.Logger
}

func NewOrderUseCase(
	orderRepository order.Repository,
	bookService book.Service,
	orderDetailRepository orderdetail.Repository,
	txManager txmanager.Repository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:       orderRepository,
		bookService:           bookService,
		orderDetailRepository: orderDetailRepository,
		txManager:             txManager,
		logger:                logger,
	}
}
