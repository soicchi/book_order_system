package orders

import (
	"log/slog"

	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CreateOrder(ctx echo.Context, dto *CreateInput) error {
	// convert dto to order entity
	order, err := order.New(dto.UserID)
	if err != nil {
		return err
	}

	// convert dto to order details entity
	orderDetails, err := ou.constructOrderDetails(dto.OrderDetails, order.ID())
	if err != nil {
		return err
	}

	order.AddOrderDetails(orderDetails)
	order.CalculateTotalPrice()

	bookIDToQuantity := ou.toQuantityMapForOrder(orderDetails)
	bookIDs := orderDetails.BookIDs()

	// update book stock
	books, err := ou.bookRepository.FindByIDs(ctx, bookIDs)
	if err != nil {
		return err
	}

	if err := books.AdjustStocks(bookIDToQuantity); err != nil {
		return err
	}

	// manage transaction
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		if err := ou.bookRepository.BulkUpdate(ctx, books); err != nil {
			return err
		}

		if err := ou.orderRepository.Create(ctx, order); err != nil {
			return err
		}

		ou.logger.Info(
			"create order success",
			slog.Any("order_id", order.ID()), slog.Any("total_price", order.TotalPrice()),
		)

		return ou.orderDetailRepository.BulkCreate(ctx, orderDetails, order.ID())
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

func (ou *OrderUseCase) toQuantityMapForOrder(orderDetails orderdetail.OrderDetails) map[uuid.UUID]int {
	bookIDToQuantity := make(map[uuid.UUID]int, len(orderDetails))
	for _, od := range orderDetails {
		bookIDToQuantity[od.BookID()] = -od.Quantity()
	}
	return bookIDToQuantity
}
