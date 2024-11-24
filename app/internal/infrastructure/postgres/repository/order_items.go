package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/orderItem"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderItemRepository struct{}

func NewOrderItemRepository() *OrderItemRepository {
	return &OrderItemRepository{}
}

func (r *OrderItemRepository) BulkCreate(
	ctx echo.Context,
	orderItem []*orderItem.OrderItem,
	orderID uuid.UUID,
	productIDs []uuid.UUID,
) error {
	db := database.GetDB(ctx)

	orderItemModels := make([]*models.OrderItem, 0, len(orderItem))

	for i := 0; i < len(orderItem); i++ {
		orderItemModels = append(orderItemModels, &models.OrderItem{
			ID:        orderItem[i].ID(),
			Quantity:  orderItem[i].Quantity(),
			OrderID:   orderID,
			ProductID: productIDs[i],
		})
	}

	err := db.Create(&orderItemModels).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("order item already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create order item: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}
