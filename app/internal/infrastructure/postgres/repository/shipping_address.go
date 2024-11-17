package repository

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	er "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
)

type ShippingAddressRepository struct{}

func NewShippingAddressRepository() *ShippingAddressRepository {
	return &ShippingAddressRepository{}
}

func (r *ShippingAddressRepository) Create(ctx echo.Context, shippingAddress *entity.ShippingAddress) error {
	db := database.GetDB(ctx)

	result := db.Create(&models.ShippingAddress{
		ID:         shippingAddress.ID(),
		Prefecture: shippingAddress.Prefecture(),
		City:       shippingAddress.City(),
		State:      shippingAddress.State(),
		CustomerID: shippingAddress.CustomerID(),
	})
	if result.Error != nil {
		return er.NewCustomError(
			fmt.Errorf("failed to create shipping address: %w", result.Error),
			er.InternalServerError,
		)
	}

	return nil
}