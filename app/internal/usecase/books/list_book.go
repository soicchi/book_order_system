package books

import (
	"github.com/labstack/echo/v4"
)

func (bu *BookUseCase) ListBook(ctx echo.Context) ([]*ListOutput, error) {
	books, err := bu.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*ListOutput, 0, len(books))

	for _, b := range books {
		output = append(output, NewListOutput(
			b.ID(),
			b.Title(),
			b.Author(),
			b.Price(),
			b.Stock(),
			b.CreatedAt().Format("2006-01-02 15:04:05"),
			b.UpdatedAt().Format("2006-01-02 15:04:05"),
		))
	}

	return output, nil
}
