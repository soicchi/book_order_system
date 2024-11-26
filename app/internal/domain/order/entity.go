package order

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/domain/values"
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
	status       values.OrderStatus
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(userID uuid.UUID) (*Order, error) {
	status, err := values.NewOrderStatus(values.Ordered)
	if err != nil {
		return nil, err
	}

	return &Order{
		id:         uuid.New(),
		userID:     userID,
		totalPrice: 0,
		orderedAt:  time.Now(),
		status:     status,
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(id uuid.UUID, userID uuid.UUID, totalPrice float64, orderedAt time.Time, status string) (*Order, error) {
	orderStatus, err := values.ReconstructOrderStatus(status)
	if err != nil {
		return nil, err
	}

	return &Order{
		id:         id,
		userID:     userID,
		totalPrice: totalPrice,
		orderedAt:  orderedAt,
		status:     orderStatus,
	}, nil
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

func (o *Order) Status() values.OrderStatus {
	return o.status
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

func (o *Order) UpdateStatus(newStatus values.OrderStatusValue) error {
	if err := o.status.Update(newStatus); err != nil {
		return err
	}

	return nil
}
