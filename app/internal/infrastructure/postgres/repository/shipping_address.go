package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	er "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (r *ShippingAddressRepository) FetchByID(ctx echo.Context, id string) (*entity.ShippingAddress, error) {
	db := database.GetDB(ctx)

	var shippingAddress models.ShippingAddress
	result := db.Where("id = ?", id).First(&shippingAddress)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, er.NewCustomError(
			fmt.Errorf("failed to fetch shipping address by id: %w", result.Error),
			er.InternalServerError,
		)
	}

	return entity.ReconstructShippingAddress(
		shippingAddress.ID,
		shippingAddress.Prefecture,
		shippingAddress.City,
		shippingAddress.State,
		shippingAddress.CustomerID,
	), nil
}
