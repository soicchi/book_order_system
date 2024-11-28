package book

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

// Bookは書籍情報を表すEntity
// 外部の層から直接変更されないようにするためにフィールドはprivateにする
// フィールドにアクセスするためのgetterメソッドを提供する
type Book struct {
	id        uuid.UUID
	title     string
	author    string
	price     float64
	stock     int
	createdAt time.Time
	updatedAt time.Time
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(title string, author string, price float64, stock int) (*Book, error) {
	if price < 0 {
		return nil, errors.New(
			fmt.Errorf("price must be greater than or equal to 0. got: %f", price),
			errors.ValidationError,
			errors.WithField("Price"),
			errors.WithIssue(errors.LessThanZero),
		)
	}

	if stock < 0 {
		return nil, errors.New(
			fmt.Errorf("stock must be greater than or equal to 0. got: %d", stock),
			errors.ValidationError,
			errors.WithField("Stock"),
			errors.WithIssue(errors.LessThanZero),
		)
	}

	return &Book{
		id:        uuid.New(),
		title:     title,
		author:    author,
		price:     price,
		stock:     stock,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(id uuid.UUID, title, author string, price float64, stock int, createdAt, updatedAt time.Time) *Book {
	return &Book{
		id:        id,
		title:     title,
		author:    author,
		price:     price,
		stock:     stock,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (b *Book) ID() uuid.UUID {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Price() float64 {
	return b.price
}

func (b *Book) Stock() int {
	return b.stock
}

func (b *Book) CreatedAt() time.Time {
	return b.createdAt
}

func (b *Book) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Book) SetStock(quantity int) error {
	b.stock += quantity

	if b.stock < 0 {
		return errors.New(
			fmt.Errorf("stock must be greater than or equal to 0. got: %d", b.stock),
			errors.ValidationError,
			errors.WithField("Stock"),
			errors.WithIssue(errors.LessThanZero),
		)
	}

	b.updatedAt = time.Now()

	return nil
}

func (b *Book) SetAll(title, author string, price float64) error {
	if price < 0 {
		return errors.New(
			fmt.Errorf("price must be greater than or equal to 0. got: %f", price),
			errors.ValidationError,
			errors.WithField("Price"),
			errors.WithIssue(errors.LessThanZero),
		)
	}

	b.title = title
	b.author = author
	b.price = price
	b.updatedAt = time.Now()

	return nil
}

type Books []*Book

func (bs Books) IDToBook() map[uuid.UUID]*Book {
	idToBook := make(map[uuid.UUID]*Book, len(bs))
	for _, b := range bs {
		idToBook[b.id] = b
	}
	return idToBook
}

func (bs Books) AdjustStocks(adjustments map[uuid.UUID]int) error {
	idToBook := bs.IDToBook()

	for id, adjustment := range adjustments {
		book := idToBook[id]
		if err := book.SetStock(adjustment); err != nil {
			return err
		}
	}

	return nil
}
