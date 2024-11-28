package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CancelOrder(ctx echo.Context, orderID uuid.UUID) error {
	owd, err := ou.orderRepository.FindByIDWithOrderDetails(ctx, orderID)
	if err != nil {
		return err
	}

	o := owd.Order()
	ods := owd.OrderDetails()

	if o == nil {
		return errors.New(
			fmt.Errorf("order not found"),
			errors.NotFoundError,
			errors.WithField("Order"),
		)
	}

	// update status to canceled in order entity
	if err := o.SetStatus(values.Cancelled); err != nil {
		return err
	}

	if len(ods) == 0 {
		return errors.New(
			fmt.Errorf("order details not found"),
			errors.NotFoundError,
			errors.WithField("OrderDetail"),
		)
	}

	bookIDToQuantity := ods.ToQuantityMapForCancel()
	bookIDs := ods.BookIDs()

	books, err := ou.bookRepository.FindByIDs(ctx, bookIDs)
	if err != nil {
		return err
	}

	if err := books.AdjustStocks(bookIDToQuantity); err != nil {
		return err
	}

	// manage transaction
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// update status to canceled
		if err := ou.orderRepository.UpdateStatus(ctx, o); err != nil {
			return err
		}

		// update book stock
		return ou.bookRepository.BulkUpdate(ctx, books)
	})
}
