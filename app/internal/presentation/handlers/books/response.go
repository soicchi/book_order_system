package books

import (
	"github.com/soicchi/book_order_system/internal/usecase/books"
)

type BookResponse struct {
	Book *books.RetrieveOutput `json:"book"`
}

func NewBookResponse(book *books.RetrieveOutput) *BookResponse {
	return &BookResponse{
		Book: book,
	}
}

type BooksResponse struct {
	Books []*books.ListOutput `json:"books"`
}

func NewBooksResponse(books []*books.ListOutput) *BooksResponse {
	return &BooksResponse{
		Books: books,
	}
}
