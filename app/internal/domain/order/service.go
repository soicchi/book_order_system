package order

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/book"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/labstack/echo/v4"
)

type OrderService struct {
	bookRepository        book.Repository
	orderRepository       Repository
	orderDetailRepository orderdetail.Repository
}

func NewOrderService(
	bookRepo book.Repository,
	orderRepo Repository,
	orderDetailRepo orderdetail.Repository,
) *OrderService {
	return &OrderService{
		bookRepository:        bookRepo,
		orderRepository:       orderRepo,
		orderDetailRepository: orderDetailRepo,
	}
}

// UseCaseケースの責務があまりにも多い場合、下記のように処理単位でサービスとして分割する
func (s *OrderService) OrderBooks(ctx echo.Context, order *Order, orderDetails []*orderdetail.OrderDetail) error {
	books, err := s.bookRepository.FindAll(ctx)
	if err != nil {
		return err
	}

	bookIDToBook := books.IDToBook()
	for _, od := range orderDetails {
		targetBook, ok := bookIDToBook[od.BookID()]
		if !ok {
			return errors.New(
				fmt.Errorf("book not found. bookID: %s", od.BookID()),
				errors.NotFound,
			)
		}

		if !targetBook.HasStock(od.Quantity()) {
			return errors.New(
				fmt.Errorf("stock is not enough. bookID: %s, stock: %d, quantity: %d", od.BookID(), targetBook.Stock(), od.Quantity()),
				errors.InvalidRequest,
			)
		}

		targetBook.ReduceStock(od.Quantity())
	}

	// Update stock
	if err := s.bookRepository.BulkUpdate(ctx, books); err != nil {
		return err
	}

	// set order details
	order.AddOrderDetails(orderDetails)

	if err := order.CalculateTotalPrice(); err != nil {
		return err
	}

	if err := s.orderRepository.Create(ctx, order); err != nil {
		return err
	}

	return s.orderDetailRepository.BulkCreate(ctx, orderDetails, order.ID())
}
