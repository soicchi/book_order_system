package orders

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CreateOrder(ctx echo.Context, dto *CreateInput) error {
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		// construct dto to order entity
		orderEntity, err := order.New(dto.UserID, 0)
		if err != nil {
			return err
		}

		orderDetails := make(orderdetail.OrderDetails, 0, len(dto.OrderDetails))
		for _, d := range dto.OrderDetails {
			od, err := orderdetail.New(orderEntity.ID(), d.BookID, d.Quantity, d.Price)
			if err != nil {
				return err
			}

			orderDetails = append(orderDetails, od)
		}

		books, err := ou.bookRepository.FindAll(ctx)
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

			targetBook.UpdateStock(-od.Quantity())
		}

		// Update stock
		if err := ou.bookRepository.BulkUpdate(ctx, books); err != nil {
			return err
		}

		// set order details
		orderEntity.AddOrderDetails(orderDetails)

		if err := orderEntity.CalculateTotalPrice(); err != nil {
			return err
		}

		if err := ou.orderRepository.Create(ctx, orderEntity); err != nil {
			return err
		}

		return ou.orderDetailRepository.BulkCreate(ctx, orderDetails, orderEntity.ID())
	})
}
