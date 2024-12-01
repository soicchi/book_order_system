package orderdetail

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

// OrderDetailは注文詳細情報を表すEntity
// 外部の層から直接変更されないようにするためにフィールドはprivateにする
// フィールドにアクセスするためのgetterメソッドを提供する
type OrderDetail struct {
	id       uuid.UUID
	quantity int
	price    float64
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(quantity int, price float64) (*OrderDetail, error) {
	if quantity <= 0 {
		return nil, errors.New(
			fmt.Errorf("quantity must be greater than 0. got: %d", quantity),
			errors.ValidationError,
			errors.WithField("Quantity"),
			errors.WithIssue(errors.ZeroOrLess),
		)
	}

	if price < 0 {
		return nil, errors.New(
			fmt.Errorf("price must be greater than or equal to 0. got: %f", price),
			errors.ValidationError,
			errors.WithField("Price"),
			errors.WithIssue(errors.LessThanZero),
		)
	}

	return &OrderDetail{
		id:       uuid.New(),
		quantity: quantity,
		price:    price,
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(id uuid.UUID, quantity int, price float64) *OrderDetail {
	return &OrderDetail{
		id:       id,
		quantity: quantity,
		price:    price,
	}
}

func (od *OrderDetail) ID() uuid.UUID {
	return od.id
}

func (od *OrderDetail) Quantity() int {
	return od.quantity
}

func (od *OrderDetail) Price() float64 {
	return od.price
}
