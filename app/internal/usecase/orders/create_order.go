package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"

	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CreateOrder(ctx echo.Context, dto *CreateInput) error {
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// construct dto to order entity
		o, err := order.New(dto.UserID, 0)
		if err != nil {
			return err
		}

		orderDetails := make(orderdetail.OrderDetails, 0, len(dto.OrderDetails))
		for _, d := range dto.OrderDetails {
			od, err := orderdetail.New(o.ID(), d.BookID, d.Quantity, d.Price)
			if err != nil {
				return err
			}

			orderDetails = append(orderDetails, od)
		}

		// create order
		if err := ou.orderService.OrderBooks(ctx, o, orderDetails); err != nil {
			return err
		}

		return ou.orderDetailRepo.BulkCreate(ctx, orderDetails, o.ID())
	})
}
