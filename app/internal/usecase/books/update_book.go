package books

import (
	"fmt"

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
			errors.NotFoundError,
			errors.WithField("Book"),
		)
	}

	if err := b.SetAll(dto.Title, dto.Author, dto.Price); err != nil {
		return err
	}

	return bu.bookRepository.Update(ctx, b)
}
