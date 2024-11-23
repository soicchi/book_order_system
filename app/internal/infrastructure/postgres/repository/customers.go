package repository

import (
	"errors"
	"fmt"

	domain "github.com/soicchi/book_order_system/internal/domain/customer"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (r *CustomerRepository) Create(ctx echo.Context, customer *domain.Customer) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Customer{
		ID:        customer.ID(),
		Name:      customer.Name(),
		Email:     customer.Email(),
		CreatedAt: *customer.CreatedAt(),
		UpdatedAt: *customer.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("customer already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create customer: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}
