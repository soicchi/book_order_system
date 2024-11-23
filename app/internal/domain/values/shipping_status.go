package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type ShippingStatus int

const (
	Waiting ShippingStatus = iota + 1
	Shipping
	Arrived
)

func NewShippingStatus(status string) (ShippingStatus, error) {
	switch status {
	case "waiting":
		return Waiting, nil
	case "shipping":
		return Shipping, nil
	case "arrived":
		return Arrived, nil
	default:
		return 0, errors.New(
			fmt.Errorf("invalid shipping status: %s", status),
			errors.InvalidRequest,
		)
	}
}

func (ss ShippingStatus) String() string {
	switch ss {
	case Waiting:
		return "waiting"
	case Shipping:
		return "shipping"
	case Arrived:
		return "arrived"
	default:
		return "unknown"
	}
}
