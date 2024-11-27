package books

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (bu *BookUseCase) RetrieveBook(ctx echo.Context, id uuid.UUID) (*RetrieveOutput, error) {
	b, err := bu.bookRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if b == nil {
		return nil, errors.New(
			fmt.Errorf("book not found"),
			errors.NotFoundError,
			errors.WithField("Book"),
		)
	}

	return NewRetrieveOutput(
		b.ID(),
		b.Title(),
		b.Author(),
		b.Price(),
		b.Stock(),
		b.CreatedAt().Format("2006-01-02 15:04:05"),
		b.UpdatedAt().Format("2006-01-02 15:04:05"),
	), nil
}
