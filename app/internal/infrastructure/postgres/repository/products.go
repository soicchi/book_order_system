package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/product"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) Create(ctx echo.Context, product *product.Product) error {
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

func (r *ProductRepository) FetchAllByIDs(ctx echo.Context, ids []uuid.UUID) (product.Products, error) {
	db := database.GetDB(ctx)

	var productModels []models.Product
	err := db.Where("id IN (?)", ids).Find(&productModels).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ers.New(
			fmt.Errorf("products not found: %w", err),
			ers.NotFound,
			ers.WithNotFoundDetails("product_ids"),
		)
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to fetch products: %w", err),
			ers.InternalServerError,
		)
	}

	var productEntities []*product.Product
	for _, pm := range productModels {
		productEntities = append(productEntities, product.Reconstruct(
			pm.ID,
			pm.Name,
			pm.Price,
			&pm.CreatedAt,
			&pm.UpdatedAt,
		))
	}

	return productEntities, nil
}
