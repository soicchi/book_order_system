package repository

import (
	"errors"
	"fmt"

	domain "github.com/soicchi/book_order_system/internal/domain/shipping"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ShippingRepository struct{}

func NewShippingRepository() *ShippingRepository {
	return &ShippingRepository{}
}

func (r *ShippingRepository) Create(ctx echo.Context, shipping *domain.Shipping, orderID uuid.UUID) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Shipping{
		ID:        shipping.ID(),
		OrderID:   orderID,
		Address:   shipping.Address(),
		CreatedAt: *shipping.CreatedAt(),
		UpdatedAt: *shipping.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("shipping already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create shipping: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}

func (r *ShippingRepository) UpdateStatus(ctx echo.Context, shipping *domain.Shipping, orderID uuid.UUID) error {
	db := database.GetDB(ctx)

	result := db.Model(&models.Shipping{}).Where("id = ?", shipping.ID()).Update("status", shipping.Status())
	if result.Error != nil {
		return ers.New(
			fmt.Errorf("failed to update shipping status: %w", result.Error),
			ers.InternalServerError,
		)
	}

	return nil
}
