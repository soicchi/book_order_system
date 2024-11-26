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

func (os OrderStatusValue) validate() error {
	switch os {
	case Ordered, Complete, Cancelled:
		return nil
	default:
		return errors.New(
			fmt.Errorf("invalid order status: %d", os),
			errors.ValidationError,
			errors.WithField(errors.OrderStatus),
			errors.WithIssue(errors.Invalid),
		)
	}
}

type OrderStatus struct {
	value OrderStatusValue
}

func NewOrderStatus(value OrderStatusValue) (OrderStatus, error) {
	if err := value.validate(); err != nil {
		return OrderStatus{}, err
	}

	return OrderStatus{
		value: value,
	}, nil
}

func ReconstructOrderStatus(value string) (OrderStatus, error) {
	switch value {
	case "ordered":
		return NewOrderStatus(Ordered)
	case "complete":
		return NewOrderStatus(Complete)
	case "cancelled":
		return NewOrderStatus(Cancelled)
	default:
		return OrderStatus{}, errors.New(
			fmt.Errorf("invalid order status: %s", value),
			errors.ValidationError,
			errors.WithField(errors.OrderStatus),
			errors.WithIssue(errors.Invalid),
		)
	}
}

func (os OrderStatus) Value() OrderStatusValue {
	return os.value
}

func (os OrderStatus) Update(value OrderStatusValue) error {
	if err := value.validate(); err != nil {
		return err
	}

	os.value = value
	return nil
}
