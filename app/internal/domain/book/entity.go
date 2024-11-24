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
			errors.InvalidRequest,
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

func new(id uuid.UUID, title, author string, price float64, stock int, createdAt, updatedAt time.Time) *Book {
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