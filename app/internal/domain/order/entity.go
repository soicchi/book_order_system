package order

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
)

// Orderは注文情報を表すEntity
// 外部の層から直接変更されないようにするためにフィールドはprivateにする
// フィールドにアクセスするためのgetterメソッドを提供する
type Order struct {
	id           uuid.UUID
	userID       uuid.UUID
	totalPrice   float64
	orderedAt    time.Time
	orderDetails []*orderdetail.OrderDetail
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(userID uuid.UUID, totalPrice float64) (*Order, error) {
	if totalPrice < 0 {
		return nil, errors.New(
			fmt.Errorf("total price must be greater than or equal to 0. got: %f", totalPrice),
			errors.InvalidRequest,
		)
	}

	return &Order{
		id:         uuid.New(),
		userID:     userID,
		totalPrice: totalPrice,
		orderedAt:  time.Now(),
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(id uuid.UUID, userID uuid.UUID, totalPrice float64, orderedAt time.Time) *Order {
	return &Order{
		id:         id,
		userID:     userID,
		totalPrice: totalPrice,
		orderedAt:  orderedAt,
	}
}

func new(id uuid.UUID, userID uuid.UUID, totalPrice float64, orderedAt time.Time) *Order {
	return &Order{
		id:         id,
		userID:     userID,
		totalPrice: totalPrice,
		orderedAt:  orderedAt,
	}
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) UserID() uuid.UUID {
	return o.userID
}

func (o *Order) TotalPrice() float64 {
	return o.totalPrice
}

func (o *Order) OrderedAt() time.Time {
	return o.orderedAt
}
