package repository

import (
	"errors"
	"fmt"

	domain "github.com/soicchi/book_order_system/internal/domain/orderItem"
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

func (r *OrderItemRepository) Create(ctx echo.Context, orderItem *domain.OrderItem, orderID, productID uuid.UUID) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.OrderItem{
		ID:        orderItem.ID(),
		Quantity:  orderItem.Quantity(),
		OrderID:   orderID,
		ProductID: productID,
	}).Error
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
