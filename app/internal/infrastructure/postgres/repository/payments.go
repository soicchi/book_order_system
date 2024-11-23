package repository

import (
	"errors"
	"fmt"

	domain "github.com/soicchi/book_order_system/internal/domain/payment"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentRepository struct{}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) Create(ctx echo.Context, payment *domain.Payment, orderID uuid.UUID) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Payment{
		ID:        payment.ID(),
		Amount:    payment.Amount(),
		OrderID:   orderID,
		CreatedAt: *payment.CreatedAt(),
		UpdatedAt: *payment.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("payment already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create payment: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}

func (r *PaymentRepository) UpdateStatus(ctx echo.Context, payment *domain.Payment, orderID uuid.UUID) error {
	db := database.GetDB(ctx)

	result := db.Model(&models.Payment{}).Where("id = ?", payment.ID()).Update("status", payment.Status())
	if result.Error != nil {
		return ers.New(
			fmt.Errorf("failed to update payment status: %w", result.Error),
			ers.InternalServerError,
		)
	}

	if result.RowsAffected == 0 {
		return ers.New(
			fmt.Errorf("payment not found: %w", result.Error),
			ers.NotFound,
			ers.WithNotFoundDetails("paymentID"),
		)
	}

	return nil
}
