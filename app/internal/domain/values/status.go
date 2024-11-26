package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type orderStatus int

const (
	Ordered orderStatus = iota + 1
	Complete
	Cancelled
)

func (os orderStatus) String() string {
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

func (os orderStatus) validate() error {
	switch os {
	case Ordered, Complete, Cancelled:
		return nil
	default:
		return errors.New(
			fmt.Errorf("invalid order status: %d", os),
			errors.InvalidRequest,
		)
	}
}

type OrderStatus struct {
	value orderStatus
}

func NewOrderStatus(value orderStatus) (OrderStatus, error) {
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
			errors.InvalidRequest,
		)
	}
}
