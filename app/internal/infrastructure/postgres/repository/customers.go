package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/values"
	er "github.com/soicchi/book_order_system/internal/errors"
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

func (r *CustomerRepository) FetchByID(ctx echo.Context, id string) (*entity.Customer, error) {
	db := database.GetDB(ctx)

	var customer models.Customer
	result := db.Where("id = ?", id).First(&customer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, er.NewCustomError(
			fmt.Errorf("customer not found: %w", result.Error),
			er.NotFound,
			er.WithNotFoundDetails("customer_id"),
		)
	}

	if result.Error != nil {
		return nil, er.NewCustomError(
			fmt.Errorf("failed to fetch customer: %w", result.Error),
			er.InternalServerError,
		)
	}

	return entity.ReconstructCustomer(
		customer.ID,
		customer.Name,
		customer.Email,
		values.Password(customer.Password),
		&customer.CreatedAt,
		&customer.UpdatedAt,
	), nil
}
