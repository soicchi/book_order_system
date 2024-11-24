package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderDetailRepository struct{}

func NewOrderDetailRepository() *OrderDetailRepository {
	return &OrderDetailRepository{}
}

func (r *OrderDetailRepository) BulkCreate(ctx echo.Context, orderDetails []*orderdetail.OrderDetail, orderID uuid.UUID) error {
	db := database.GetDB(ctx)

	var ods []models.OrderDetail
	for _, od := range orderDetails {
		ods = append(ods, models.OrderDetail{
			ID:       od.ID(),
			OrderID:  orderID,
			BookID:   od.BookID(),
			Quantity: od.Quantity(),
			Price:    od.Price(),
		})
	}

	err := db.Create(&ods).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("order detail already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create order detail: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}
