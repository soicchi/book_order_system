package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	errorsPkg "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (r *CustomerRepository) Create(ctx echo.Context, customer *entity.Customer) error {
	db := database.GetDB(ctx)

	result := db.Create(&models.Customer{
		ID:       customer.ID(),
		Name:     customer.Name(),
		Email:    customer.Email(),
		Password: customer.Password().Value(),
	})
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return errorsPkg.NewCustomError(
			fmt.Errorf("email already exists: %w", result.Error),
			errorsPkg.AlreadyExist,
		)
	}
	if result.Error != nil {
		return errorsPkg.NewCustomError(
			fmt.Errorf("failed to create customer: %w", result.Error),
			errorsPkg.InternalServerError,
		)
	}

	return nil
}
