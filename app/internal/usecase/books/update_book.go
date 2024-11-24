package books

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/book"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/labstack/echo/v4"
)

func (bu *BookUseCase) UpdateBook(ctx echo.Context, dto *UpdateInput) error {
	b, err := bu.bookRepository.FindByID(ctx, dto.ID)
	if err != nil {
		return err
	}

	if b == nil {
		return errors.New(
			fmt.Errorf("book not found"),
			errors.NotFound,
		)
	}

	// Reconstruct関数を利用して、書籍情報を更新する
	// 既存のEntityを更新する場合は、Reconstruct関数を利用する
	updatedBook := book.Reconstruct(
		b.ID(),
		dto.Title,
		dto.Author,
		dto.Price,
		dto.Stock,
		b.CreatedAt(),
		time.Now(),
	)

	return bu.bookRepository.Update(ctx, updatedBook)
}
