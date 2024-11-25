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
	orderDetails orderdetail.OrderDetails
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(userID uuid.UUID) (*Order, error) {
	return &Order{
		id:         uuid.New(),
		userID:     userID,
		totalPrice: 0,
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

func (o *Order) OrderDetails() orderdetail.OrderDetails {
	return o.orderDetails
}

func (o *Order) AddOrderDetails(orderDetails orderdetail.OrderDetails) {
	o.orderDetails = append(o.orderDetails, orderDetails...)
}

func (o *Order) CalculateTotalPrice() error {
	var totalPrice float64
	for _, od := range o.orderDetails {
		totalPrice += od.Price()
	}

	if totalPrice < 0 {
		return errors.New(
			fmt.Errorf("total price must be greater than or equal to 0. got: %f", totalPrice),
			errors.InvalidRequest,
		)
	}

	o.totalPrice = totalPrice

	return nil
}
