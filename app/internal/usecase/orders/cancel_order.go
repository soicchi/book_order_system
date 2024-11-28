package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CancelOrder(ctx echo.Context, orderID uuid.UUID) error {
	// Fetch order and order details by once
	order, err := ou.orderRepository.FindByIDWithOrderDetails(ctx, orderID)
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

	// update status to canceled in order entity
	if err := order.ChangeStatus(values.Cancelled); err != nil {
		return err
	}

	orderDetails := order.OrderDetails()
	if len(orderDetails) == 0 {
		return errors.New(
			fmt.Errorf("order details not found"),
			errors.NotFoundError,
			errors.WithField("OrderDetail"),
		)
	}

	bookIDs := orderDetails.BookIDs()

	books, err := ou.bookRepository.FindByIDs(ctx, bookIDs)
	if err != nil {
		return err
	}

	bookIDToQuantity := ou.toQuantityMapForCancel(orderDetails)
	if err := books.AdjustStocks(bookIDToQuantity); err != nil {
		return err
	}

	// manage transaction
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// update status to canceled
		if err := ou.orderRepository.UpdateStatus(ctx, order); err != nil {
			return err
		}

		// update book stock
		return ou.bookRepository.BulkUpdate(ctx, books)
	})
}

func (ou *OrderUseCase) toQuantityMapForCancel(orderDetails orderdetail.OrderDetails) map[uuid.UUID]int {
	bookIDToQuantity := make(map[uuid.UUID]int, len(orderDetails))
	for _, od := range orderDetails {
		bookIDToQuantity[od.BookID()] = od.Quantity()
	}
	return bookIDToQuantity
}
