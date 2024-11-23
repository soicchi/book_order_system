package repository

import (
	"errors"
	"fmt"

	domain "github.com/soicchi/book_order_system/internal/domain/order"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(ctx echo.Context, order *domain.Order, customerID uuid.UUID) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Order{
		ID:          order.ID(),
		CustomerID:  customerID,
		TotalAmount: order.TotalAmount(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("order already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create order: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}
