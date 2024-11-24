package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderItem"
	"github.com/soicchi/book_order_system/internal/domain/product"
	"github.com/soicchi/book_order_system/internal/domain/transactionManager"
	"github.com/soicchi/book_order_system/internal/logging"
)

type OrderUseCase struct {
	orderRepo     order.Repository
	orderItemRepo orderItem.Repository
	productRepo   product.Repository
	txManager     transactionManager.Repository
	logger        logging.Logger
}

func NewOrderUseCase(
	orderRepo order.Repository,
	orderItemRepo orderItem.Repository,
	productRepo product.Repository,
	txManager transactionManager.Repository,
	logger logging.Logger,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		productRepo:   productRepo,
		txManager:     txManager,
		logger:        logger,
	}
}
