package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CancelOrder(ctx echo.Context, orderID uuid.UUID) error {
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		o, err := ou.orderRepository.FindByID(ctx, orderID)
		if err != nil {
			return err
		}

		if o == nil {
			return errors.New(fmt.Errorf("order not found"), errors.NotFound)
		}

		// update status to canceled in order entity
		if err := o.UpdateStatus(values.Cancelled); err != nil {
			return err
		}

		// update status to canceled in order repository
		if err := ou.orderRepository.UpdateStatus(ctx, o); err != nil {
			return err
		}

		ods, err := ou.orderDetailRepository.FindByOrderID(ctx, orderID)
		if err != nil {
			return err
		}

		if ods == nil {
			return errors.New(fmt.Errorf("order details not found"), errors.NotFound)
		}

		bookIDToQuantity := ods.AdjustmentInCancel()

		// update book stock
		if err := ou.bookService.UpdateStock(ctx, bookIDToQuantity); err != nil {
			return err
		}

		return nil
	})
}
