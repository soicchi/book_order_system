package repository

import (
	"errors"
	"fmt"

	domain "github.com/soicchi/book_order_system/internal/domain/product"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) Create(ctx echo.Context, product *domain.Product) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Product{
		ID:        product.ID(),
		Name:      product.Name(),
		Price:     product.Price(),
		CreatedAt: *product.CreatedAt(),
		UpdatedAt: *product.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("product already exists: %w", err),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create product: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}
