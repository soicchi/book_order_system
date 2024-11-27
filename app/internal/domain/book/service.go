package book

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Service struct {
	bookRepository Repository
}

func NewService(bookRepository Repository) *Service {
	return &Service{
		bookRepository: bookRepository,
	}
}

// 下記のように他のケースでも利用（注文のキャンセル処理など）できる一連の操作はドメインサービスを用いて再利用できるようにする。
func (bs *Service) UpdateStock(ctx echo.Context, adjustments map[uuid.UUID]int) error {
	bookIDs := make([]uuid.UUID, 0, len(adjustments))

	for bookID := range adjustments {
		bookIDs = append(bookIDs, bookID)
	}

	books, err := bs.bookRepository.FindByIDs(ctx, bookIDs)
	if err != nil {
		return err
	}

	bookIDToBook := books.IDToBook()

	for bookID, adjustment := range adjustments {
		book := bookIDToBook[bookID]
		if err := book.UpdateStock(adjustment); err != nil {
			return err
		}
	}

	return bs.bookRepository.BulkUpdate(ctx, books)
}
