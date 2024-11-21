package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/domain/values"
	er "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (r *CustomerRepository) Create(ctx echo.Context, customer *customer.Customer) error {
	db := database.GetDB(ctx)

	result := db.Create(&models.Customer{
		ID:       customer.ID(),
		Name:     customer.Name(),
		Email:    customer.Email(),
		Password: customer.Password().String(),
	})
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return er.NewCustomError(
			fmt.Errorf("email already exists: %w", result.Error),
			er.AlreadyExist,
		)
	}
	if result.Error != nil {
		return er.NewCustomError(
			fmt.Errorf("failed to create customer: %w", result.Error),
			er.InternalServerError,
		)
	}

	return nil
}

func (r *CustomerRepository) FetchByID(ctx echo.Context, id uuid.UUID) (*customer.Customer, error) {
	db := database.GetDB(ctx)

	var customerModel models.Customer
	result := db.Where("id = ?", id).First(&customerModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, er.NewCustomError(
			fmt.Errorf("failed to fetch customer: %w", result.Error),
			er.InternalServerError,
		)
	}

	return customer.Reconstruct(
		customerModel.ID,
		customerModel.Name,
		customerModel.Email,
		values.Password(customerModel.Password),
		&customerModel.CreatedAt,
		&customerModel.UpdatedAt,
	), nil
}
