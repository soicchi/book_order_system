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
			ers.AlreadyExistError,
			ers.WithField("OrderDetail"),
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create order detail: %w", err),
			ers.UnexpectedError,
		)
	}

	return nil
}

func (r *OrderDetailRepository) FindByOrderID(ctx echo.Context, orderID uuid.UUID) (orderdetail.OrderDetails, error) {
	db := database.GetDB(ctx)

	var ods []models.OrderDetail
	err := db.Where("order_id = ?", orderID).Find(&ods).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find order details by order id: %w", err),
			ers.UnexpectedError,
		)
	}

	var orderDetails orderdetail.OrderDetails
	for _, od := range ods {
		orderDetails = append(orderDetails, orderdetail.Reconstruct(
			od.ID,
			od.BookID,
			od.BookID,
			od.Quantity,
			od.Price,
		))
	}

	return orderDetails, nil
}
