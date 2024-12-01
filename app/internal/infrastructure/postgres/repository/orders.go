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

func (r *OrderRepository) Create(ctx echo.Context, order *order.Order, bookID uuid.UUID) error {
	db := database.GetDB(ctx)
	orderDetailModels := r.toOrderDetailModels(order.OrderDetails(), order.ID(), bookID)
	orderModel := r.toOrderModel(order, orderDetailModels)

	err := db.Create(&orderModel).Error

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

func (r *OrderRepository) FindByID(ctx echo.Context, orderID uuid.UUID) (*order.Order, error) {
	db := database.GetDB(ctx)

	var o models.Order
	err := db.Preload("OrderDetails").Where("id = ?", orderID).First(&o).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find order: %w", err),
			ers.UnexpectedError,
		)
	}

	return order.Reconstruct(
		o.ID,
		o.TotalPrice,
		o.OrderedAt,
		o.Status,
		r.reconstructOrderDetails(o.OrderDetails),
	), nil
}

func (r *OrderRepository) toOrderModel(order *order.Order, orderDetailModels []models.OrderDetail) models.Order {
	return models.Order{
		ID:           order.ID(),
		TotalPrice:   order.TotalPrice(),
		OrderedAt:    order.OrderedAt(),
		Status:       order.Status().Value().String(),
		OrderDetails: orderDetailModels,
	}
}

func (r *OrderRepository) reconstructOrderDetails(orderDetails []models.OrderDetail) []*orderdetail.OrderDetail {
	orderDetailEntities := make([]*orderdetail.OrderDetail, 0, len(orderDetails))
	for _, od := range orderDetails {
		orderDetailEntities = append(
			orderDetailEntities,
			orderdetail.Reconstruct(od.ID, od.Quantity, od.Price),
		)
	}
	return orderDetailEntities
}

func (r *OrderRepository) toOrderDetailModels(
	orderDetails []*orderdetail.OrderDetail,
	orderID uuid.UUID,
	bookID uuid.UUID,
) []models.OrderDetail {
	orderDetailModels := make([]models.OrderDetail, 0, len(orderDetails))
	for _, od := range orderDetails {
		orderDetailModels = append(orderDetailModels, models.OrderDetail{
			ID:       od.ID(),
			OrderID:  orderID,
			BookID:   bookID,
			Quantity: od.Quantity(),
			Price:    od.Price(),
		})
	}
	return orderDetailModels
}
