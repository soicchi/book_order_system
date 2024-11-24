package orders

import (
	"fmt"

	orderDomain "github.com/soicchi/book_order_system/internal/domain/order"
	orderItemDomain "github.com/soicchi/book_order_system/internal/domain/orderItem"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (u *OrderUseCase) CreateOrder(ctx echo.Context, input *CreateInput) error {
	orderItems, productIDs, err := u.prepareOrderItems(ctx, input)
	if err != nil {
		return err
	}

	products, err := u.productRepo.FetchAllByIDs(ctx, productIDs)
	if err != nil {
		return err
	}

	totalAmount := products.CalculateTotalPrice()

	order, err := orderDomain.New(totalAmount)
	if err != nil {
		return err
	}

	customerUUID, err := uuid.Parse(input.CustomerID)
	if err != nil {
		return errors.New(
			fmt.Errorf("failed to parse customer UUID: %w", err),
			errors.InternalServerError,
		)
	}

	return u.txManager.WithTransaction(ctx, func(ctx echo.Context) error {
		if err := u.orderRepo.Create(ctx, order, customerUUID); err != nil {
			return err
		}

		if err := u.orderItemRepo.BulkCreate(ctx, orderItems, order.ID(), productIDs); err != nil {
			return err
		}

		return nil
	})
}

func (u *OrderUseCase) prepareOrderItems(ctx echo.Context, input *CreateInput) ([]*orderItemDomain.OrderItem, []uuid.UUID, error) {
	orderItems := make([]*orderItemDomain.OrderItem, 0, len(input.OrderItems))
	productIDs := make([]uuid.UUID, 0, len(input.OrderItems))

	for _, item := range input.OrderItems {
		orderItem, err := orderItemDomain.New(item.Quantity)
		if err != nil {
			return nil, nil, err
		}

		productUUID, err := uuid.Parse(item.ProductID)
		if err != nil {
			return nil, nil, errors.New(
				fmt.Errorf("failed to parse product UUID: %w", err),
				errors.InternalServerError,
			)
		}

		orderItems = append(orderItems, orderItem)
		productIDs = append(productIDs, productUUID)
	}

	return orderItems, productIDs, nil
}
