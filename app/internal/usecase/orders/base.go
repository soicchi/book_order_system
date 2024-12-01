package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/book"
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/txmanager"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderRepository order.OrderRepository
	bookRepository  book.BookRepository
	txManager       txmanager.Repository
	logger          logging.Logger
}

func NewUseCase(
	orderRepository order.OrderRepository,
	bookRepository book.BookRepository,
	txManager txmanager.Repository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
		bookRepository:  bookRepository,
		txManager:       txManager,
		logger:          logger,
	}
}
