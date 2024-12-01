package orders

// import (
// 	"github.com/google/uuid"
// )

// func (ou *OrderUseCase) adjustBookStocks(adjustments map[uuid.UUID]int) error {
// 	bookIDs := make([]uuid.UUID, 0, len(bookIDToQuantity))
// 	for bookID := range bookIDToQuantity {
// 		bookIDs = append(bookIDs, bookID)
// 	}

// 	// update book stock
// 	books, err := ou.bookRepository.FindByIDs(ctx, bookIDs)
// 	if err != nil {
// 		return err
// 	}

// 	bookIDToBook := make(map[uuid.UUID]*book.Book, len(books))
// 	for _, b := range books {
// 		bookIDToBook[b.ID()] = b
// 	}

// 	for bookID, adjustment := range adjustments {
// 		book := bookIDToBook[bookID]
// 		if err := book.SetStock(adjustment); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
