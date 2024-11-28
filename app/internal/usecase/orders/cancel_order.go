package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CancelOrder(ctx echo.Context, orderID uuid.UUID) error {
	o, err := ou.orderRepository.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

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

	ods, err := ou.orderDetailRepository.FindByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	if ods == nil {
		return errors.New(
			fmt.Errorf("order details not found"),
			errors.NotFoundError,
			errors.WithField("OrderDetail"),
		)
	}

	// manage transaction
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// update status to canceled in order repository
		if err := ou.orderRepository.UpdateStatus(ctx, o); err != nil {
			return err
		}

		bookIDToQuantity := ods.ToQuantityMapForCancel()

		// update book stock
		if err := ou.bookService.UpdateStock(ctx, bookIDToQuantity); err != nil {
			return err
		}

		return nil
	})
}
