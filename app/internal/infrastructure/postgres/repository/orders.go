package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
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
			ers.AlreadyExistError,
			ers.WithField("Order"),
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create order: %w", err),
			ers.UnexpectedError,
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
			ers.UnexpectedError,
		)
	}

	return order.Reconstruct(o.ID, o.UserID, o.TotalPrice, o.OrderedAt, o.Status), nil
}

func (r *OrderRepository) FindByIDWithOrderDetails(
	ctx echo.Context,
	id uuid.UUID,
) (*order.OrderWithDetails, error) {
	db := database.GetDB(ctx)

	var orderModel models.Order
	err := db.Where("id = ?", id).Preload("OrderDetails").First(&orderModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find order: %w", err),
			ers.UnexpectedError,
		)
	}

	o := order.Reconstruct(orderModel.ID, orderModel.UserID, orderModel.TotalPrice, orderModel.OrderedAt, orderModel.Status)
	ods := make(orderdetail.OrderDetails, 0, len(orderModel.OrderDetails))
	for _, od := range orderModel.OrderDetails {
		ods = append(ods, orderdetail.Reconstruct(od.ID, od.OrderID, od.BookID, od.Quantity, od.Price))
	}

	return order.ReconstructOrderWithDetails(o, ods), nil
}

func (r *OrderRepository) UpdateStatus(ctx echo.Context, order *order.Order) error {
	db := database.GetDB(ctx)

	result := db.Model(&models.Order{}).Where("id = ?", order.ID()).Update("status", order.Status().Value().String())
	if result.Error != nil {
		return ers.New(
			fmt.Errorf("failed to update order status: %w", result.Error),
			ers.UnexpectedError,
		)
	}

	if result.RowsAffected == 0 {
		return ers.New(
			fmt.Errorf("order not found: %w", result.Error),
			ers.NotFoundError,
			ers.WithField("Order"),
		)
	}

	return nil
}
