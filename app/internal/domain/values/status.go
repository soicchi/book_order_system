package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type OrderStatusValue int

const (
	Ordered OrderStatusValue = iota + 1
	Complete
	Cancelled
)

func (os OrderStatusValue) String() string {
	switch os {
	case Ordered:
		return "ordered"
	case Complete:
		return "complete"
	case Cancelled:
		return "cancelled"
	default:
		return ""
	}
}

type OrderStatus struct {
	value OrderStatusValue
}

func NewOrderStatus(value OrderStatusValue) (OrderStatus, error) {
	if err := validate(value); err != nil {
		return OrderStatus{}, err
	}

	return OrderStatus{
		value: value,
	}, nil
}

func validate(value OrderStatusValue) error {
	switch value {
	case Ordered, Complete, Cancelled:
		return nil
	default:
		return errors.New(
			fmt.Errorf("invalid order status: %d", value),
			errors.ValidationError,
			errors.WithField("Status"),
			errors.WithIssue(errors.Invalid),
		)
	}
}

func ReconstructOrderStatus(value string) OrderStatus {
	switch value {
	case "ordered":
		return OrderStatus{value: Ordered}
	case "complete":
		return OrderStatus{value: Complete}
	case "cancelled":
		return OrderStatus{value: Cancelled}
	default:
		return OrderStatus{}
	}
}

func (os OrderStatus) Value() OrderStatusValue {
	return os.value
}

func (os OrderStatus) Set(value OrderStatusValue) error {
	if err := validate(value); err != nil {
		return err
	}

	os.value = value
	return nil
}
