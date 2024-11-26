package books

import (
	"github.com/soicchi/book_order_system/internal/domain/book"

	"github.com/labstack/echo/v4"
)

func (bu *BookUseCase) CreateBook(ctx echo.Context, dto *CreateInput) error {
	b, err := book.New(dto.Title, dto.Author, dto.Price, dto.Stock)
	if err != nil {
		return err
	}

	return bu.bookRepository.Create(ctx, b)
}
