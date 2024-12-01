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
	totalPrice   float64
	orderedAt    time.Time
	status       values.OrderStatus
	orderDetails []*orderdetail.OrderDetail
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(orderDetails []*orderdetail.OrderDetail) (*Order, error) {
	status, err := values.NewOrderStatus(values.Ordered)
	if err != nil {
		return nil, err
	}

	if len(orderDetails) == 0 {
		return nil, errors.New(
			fmt.Errorf("order details must be greater than 0. got: %d", len(orderDetails)),
			errors.ValidationError,
			errors.WithField("OrderDetails"),
			errors.WithIssue(errors.ZeroOrLess),
		)
	}

	var totalPrice float64
	for _, od := range orderDetails {
		totalPrice += od.Price()
	}

	if totalPrice < 0 {
		return nil, errors.New(
			fmt.Errorf("total price must be greater than or equal to 0. got: %f", totalPrice),
			errors.ValidationError,
			errors.WithField("TotalPrice"),
			errors.WithIssue(errors.LessThanZero),
		)
	}

	return &Order{
		id:           uuid.New(),
		totalPrice:   0,
		orderedAt:    time.Now(),
		status:       status,
		orderDetails: orderDetails,
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(
	id uuid.UUID,
	totalPrice float64,
	orderedAt time.Time,
	status string,
	orderDetails []*orderdetail.OrderDetail,
) *Order {
	orderStatus := values.ReconstructOrderStatus(status)

	return &Order{
		id:           id,
		totalPrice:   totalPrice,
		orderedAt:    orderedAt,
		status:       orderStatus,
		orderDetails: orderDetails,
	}
}

func (o *Order) ID() uuid.UUID {
	return o.id
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

func (o *Order) OrderDetails() []*orderdetail.OrderDetail {
	return o.orderDetails
}

func (o *Order) SetStatus(newStatus values.OrderStatusValue) error {
	if err := o.status.Set(newStatus); err != nil {
		return err
	}

	return nil
}
