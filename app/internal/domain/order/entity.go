package order

import (
	"time"

	"github.com/soicchi/book_order_system/internal/domain/orderdetail"
	"github.com/soicchi/book_order_system/internal/domain/values"

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
	status       values.OrderStatus
	orderDetails orderdetail.OrderDetails
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
func Reconstruct(id uuid.UUID, userID uuid.UUID, totalPrice float64, orderedAt time.Time, status string) *Order {
	orderStatus := values.ReconstructOrderStatus(status)

	return &Order{
		id:         id,
		userID:     userID,
		totalPrice: totalPrice,
		orderedAt:  orderedAt,
		status:     orderStatus,
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

func (o *Order) Status() values.OrderStatus {
	return o.status
}

func (o *Order) OrderDetails() orderdetail.OrderDetails {
	return o.orderDetails
}

func (o *Order) AddOrderDetails(od orderdetail.OrderDetails) {
	o.orderDetails = append(o.orderDetails, od...)
}

func (o *Order) CalculateTotalPrice() {
	var totalPrice float64
	for _, od := range o.orderDetails {
		totalPrice += od.Price()
	}

	o.totalPrice = totalPrice
}

func (o *Order) ChangeStatus(newStatus values.OrderStatusValue) error {
	if err := o.status.Set(newStatus); err != nil {
		return err
	}

	return nil
}
