package orders

import (
	"github.com/soicchi/book_order_system/internal/domain/book"
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ou *OrderUseCase) CreateOrder(ctx echo.Context, dto *CreateInput) error {
	// convert dto to order details entity
	orderDetails, err := ou.constructOrderDetails(dto.OrderDetails)
	if err != nil {
		return err
	}

	// convert dto to order entity
	order, err := order.New(orderDetails)
	if err != nil {
		return err
	}

	bookIDs := make([]uuid.UUID, 0, len(orderDetails))
	for _, d := range dto.OrderDetails {
		bookIDs = append(bookIDs, d.BookID)
	}

	books, err := ou.bookRepository.FindByIDs(ctx, bookIDs)
	if err != nil {
		return err
	}

	bookIDToBook := make(map[uuid.UUID]*book.Book, len(books))
	for _, b := range books {
		bookIDToBook[b.ID()] = b
	}

	// update book stock
	for _, d := range dto.OrderDetails {
		book := bookIDToBook[d.BookID]
		if err := book.SetStock(d.Quantity); err != nil {
			return err
		}
	}

	// manage transaction
	return ou.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		if err := ou.bookRepository.BulkUpdate(ctx, books); err != nil {
			return err
		}

		return ou.orderRepository.Create(ctx, order, dto.UserID)
	})
}

func (ou *OrderUseCase) constructOrderDetails(dto []*CreateDetailInput) ([]*orderdetail.OrderDetail, error) {
	ods := make([]*orderdetail.OrderDetail, 0, len(dto))
	for _, d := range dto {
		od, err := orderdetail.New(d.Quantity, d.Price)
		if err != nil {
			return nil, err
		}

		ods = append(ods, od)
	}

	return ods, nil
}
