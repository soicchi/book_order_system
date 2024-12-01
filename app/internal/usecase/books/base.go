package books

import (
	"github.com/soicchi/book_order_system/internal/domain/book"
	"github.com/soicchi/book_order_system/internal/logging"
)

type BookUseCase struct {
	bookRepository book.BookRepository
	logger         logging.Logger
}

func NewUseCase(bookRepository book.BookRepository, logger logging.Logger) *BookUseCase {
	return &BookUseCase{
		bookRepository: bookRepository,
		logger:         logger,
	}
}
