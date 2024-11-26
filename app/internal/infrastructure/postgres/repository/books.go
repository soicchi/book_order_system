package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/book"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Create(ctx echo.Context, book *book.Book) error {
	db := database.GetDB(ctx)

	err := db.Create(&models.Book{
		ID:        book.ID(),
		Title:     book.Title(),
		Author:    book.Author(),
		Price:     book.Price(),
		CreatedAt: book.CreatedAt(),
		UpdatedAt: book.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("book with title %s already exists", book.Title()),
			ers.AlreadyExistError,
			ers.WithField(ers.Book),
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create book: %w", err),
			ers.UnexpectedError,
		)
	}

	return nil
}

func (r *BookRepository) FindByID(ctx echo.Context, id uuid.UUID) (*book.Book, error) {
	db := database.GetDB(ctx)

	var b models.Book
	err := db.Where("id = ?", id).First(&b).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find book by id: %w", err),
			ers.UnexpectedError,
		)
	}

	return book.Reconstruct(
		b.ID,
		b.Title,
		b.Author,
		b.Price,
		b.Stock,
		b.CreatedAt,
		b.UpdatedAt,
	), nil
}

func (r *BookRepository) FindAll(ctx echo.Context) (book.Books, error) {
	db := database.GetDB(ctx)

	var bs []models.Book
	result := db.Find(&bs)
	if result.Error != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find all books: %w", result.Error),
			ers.UnexpectedError,
		)
	}

	books := make(book.Books, 0, len(bs))
	for _, b := range bs {
		books = append(books, book.Reconstruct(
			b.ID,
			b.Title,
			b.Author,
			b.Price,
			b.Stock,
			b.CreatedAt,
			b.UpdatedAt,
		))
	}

	return books, nil
}

func (r *BookRepository) FindByIDs(ctx echo.Context, ids []uuid.UUID) ([]*book.Book, error) {
	db := database.GetDB(ctx)

	var bs []models.Book
	result := db.Find(&bs, ids)
	if result.Error != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find all books: %w", result.Error),
			ers.UnexpectedError,
		)
	}

	if len(bs) != len(ids) {
		return nil, ers.New(
			fmt.Errorf("failed to find all books by ids: %w", gorm.ErrRecordNotFound),
			ers.NotFoundError,
			ers.WithField(ers.Book),
		)
	}

	var books book.Books
	for _, b := range bs {
		books = append(books, book.Reconstruct(
			b.ID,
			b.Title,
			b.Author,
			b.Price,
			b.Stock,
			b.CreatedAt,
			b.UpdatedAt,
		))
	}

	return books, nil
}

func (r *BookRepository) Update(ctx echo.Context, book *book.Book) error {
	db := database.GetDB(ctx)

	err := db.Model(&models.Book{}).Where("id = ?", book.ID()).Updates(models.Book{
		Title:     book.Title(),
		Author:    book.Author(),
		Price:     book.Price(),
		Stock:     book.Stock(),
		UpdatedAt: book.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ers.New(
			fmt.Errorf("book with id %s not found", book.ID()),
			ers.NotFoundError,
			ers.WithField(ers.Book),
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to update book: %w", err),
			ers.UnexpectedError,
		)
	}

	return nil
}

func (r *BookRepository) BulkUpdate(ctx echo.Context, books []*book.Book) error {
	db := database.GetDB(ctx)

	var bs []models.Book
	for _, b := range books {
		bs = append(bs, models.Book{
			ID:        b.ID(),
			Title:     b.Title(),
			Author:    b.Author(),
			Price:     b.Price(),
			Stock:     b.Stock(),
			UpdatedAt: b.UpdatedAt(),
		})
	}

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "author", "price", "stock", "updated_at"}),
	}).Create(&bs).Error; err != nil {
		return ers.New(
			fmt.Errorf("failed to bulk update books: %w", err),
			ers.UnexpectedError,
		)
	}

	return nil
}
