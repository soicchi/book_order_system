package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type ShippingMethod int

const (
	Standard ShippingMethod = iota + 1
	Express
)

func NewShippingMethod(method string) (ShippingMethod, error) {
	switch method {
	case "standard":
		return Standard, nil
	case "express":
		return Express, nil
	default:
		return 0, errors.New(
			fmt.Errorf("invalid shipping method: %s", method),
			errors.InvalidRequest,
		)
	}
}

func (sm ShippingMethod) String() string {
	switch sm {
	case Standard:
		return "standard"
	case Express:
		return "express"
	default:
		return "unknown"
	}
}
