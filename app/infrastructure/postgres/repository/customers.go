package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/domain/entity"
	"github.com/soicchi/book_order_system/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	dbConn DBConnector
}

func NewCustomerRepository(dbConn DBConnector) *CustomerRepository {
	return &CustomerRepository{
		dbConn: dbConn,
	}
}

func (r *CustomerRepository) Create(ctx echo.Context, customer *entity.Customer) error {
	db := r.dbConn.GetDB(ctx)

	result := db.Create(&models.Customer{
		ID:       customer.ID(),
		Name:     customer.Name(),
		Email:    customer.Email(),
		Password: customer.Password().Value(),
	})
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return fmt.Errorf("email %s is already registered", customer.Email())
	}
	if result.Error != nil {
		return fmt.Errorf("error creating customer: %w", result.Error)
	}

	return nil
}
