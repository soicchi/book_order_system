package orders

import (
	"log/slog"

	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CreateOrder(ctx echo.Context, dto *CreateInput) error {
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// convert dto to order entity
		o, err := order.New(dto.UserID)
		if err != nil {
			return err
		}

		// convert dto to order details entity
		ods, err := ou.constructOrderDetails(dto.OrderDetails, o.ID())
		if err != nil {
			return err
		}

		bookIDToQuantity := ods.ToQuantityMapForOrder()

		// update book stock
		if err := ou.bookService.UpdateStock(ctx, bookIDToQuantity); err != nil {
			return err
		}

		// set order details
		o.AddOrderDetails(ods)

		if err := o.CalculateTotalPrice(); err != nil {
			return err
		}

		if err := ou.orderRepository.Create(ctx, o); err != nil {
			return err
		}

		ou.logger.Info("create order success", slog.Any("order_id", o.ID()), slog.Any("total_price", o.TotalPrice()))

		return ou.orderDetailRepository.BulkCreate(ctx, ods, o.ID())
	})
}

func (ou *OrderUseCase) constructOrderDetails(dto []*CreateDetailInput, orderID uuid.UUID) (orderdetail.OrderDetails, error) {
	ods := make(orderdetail.OrderDetails, 0, len(dto))
	for _, d := range dto {
		od, err := orderdetail.New(orderID, d.BookID, d.Quantity, d.Price)
		if err != nil {
			return nil, err
		}

		ods = append(ods, od)
	}

	return ods, nil
}
