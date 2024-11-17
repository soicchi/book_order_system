package repository

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	er "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(ctx echo.Context, order *entity.Order) error {
	db := database.GetDB(ctx)

	result := db.Create(&models.Order{
		ID:                order.ID(),
		CustomerID:        order.CustomerID(),
		ShippingAddressID: order.ShippingAddressID(),
	})
	if result.Error != nil {
		return er.NewCustomError(
			fmt.Errorf("failed to create order: %w", result.Error),
			er.InternalServerError,
		)
	}

	return nil
}
