package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/book"
	orderDomain "github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CancelOrder(ctx echo.Context, dto *CancelInput) error {
	// Fetch order and order details by once
	order, err := ou.orderRepository.FindByID(ctx, dto.OrderID)
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New(
			fmt.Errorf("order not found"),
			errors.NotFoundError,
			errors.WithField("Order"),
		)
	}

	// Can't cancel if order status is not ordered
	if order.Status().Value() != values.Ordered {
		return errors.New(
			fmt.Errorf("order status is not ordered. got: %s", order.Status().Value()),
			errors.ValidationError,
			errors.WithField("OrderStatus"),
			errors.WithIssue(errors.Invalid),
		)
	}

	orderDetails := order.OrderDetails()

	newOrder, err := orderDomain.New(orderDetails)
	if err != nil {
		return err
	}

	bookIDs := make([]uuid.UUID, 0, len(dto.Details))
	for _, detail := range dto.Details {
		bookIDs = append(bookIDs, detail.BookID)
	}

	books, err := ou.bookRepository.FindByIDs(ctx, bookIDs)
	if err != nil {
		return err
	}

	bookIDToBook := make(map[uuid.UUID]*book.Book, len(books))
	for _, b := range books {
		bookIDToBook[b.ID()] = b
	}

	// set book stock
	for _, detail := range dto.Details {
		book := bookIDToBook[detail.BookID]
		if err := book.SetStock(detail.Quantity); err != nil {
			return err
		}
	}

	// manage transaction
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// update status to canceled
		if err := ou.orderRepository.Create(ctx, newOrder, dto.UserID); err != nil {
			return err
		}

		// update book stock
		return ou.bookRepository.BulkUpdate(ctx, books)
	})
}
