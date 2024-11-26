package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/order"
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

func (r *OrderRepository) Create(ctx echo.Context, order *order.Order) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Order{
		ID:         order.ID(),
		UserID:     order.UserID(),
		TotalPrice: order.TotalPrice(),
		OrderedAt:  order.OrderedAt(),
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

func (r *OrderRepository) FindByID(ctx echo.Context, id uuid.UUID) (*order.Order, error) {
	db := database.GetDB(ctx)

	var o models.Order
	err := db.Where("id = ?", id).First(&o).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find order: %w", err),
			ers.InternalServerError,
		)
	}

	return order.Reconstruct(o.ID, o.UserID, o.TotalPrice, o.OrderedAt, o.Status)
}

func (r *OrderRepository) UpdateStatus(ctx echo.Context, order *order.Order) error {
	db := database.GetDB(ctx)

	result := db.Model(&models.Order{}).Where("id = ?", order.ID()).Update("status", order.Status().Value().String())
	if result.Error != nil {
		return ers.New(
			fmt.Errorf("failed to update order status: %w", result.Error),
			ers.InternalServerError,
		)
	}

	if result.RowsAffected == 0 {
		return ers.New(
			fmt.Errorf("order not found: %w", result.Error),
			ers.NotFound,
		)
	}

	return nil
}
