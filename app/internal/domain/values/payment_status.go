package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type PaymentStatus int

const (
	Pending PaymentStatus = iota + 1
	Failed
	Success
)

func NewPaymentStatus(status string) (PaymentStatus, error) {
	switch status {
	case "pending":
		return Pending, nil
	case "failed":
		return Failed, nil
	case "success":
		return Success, nil
	default:
		return 0, errors.New(
			fmt.Errorf("invalid payment status: %s", status),
			errors.InvalidRequest,
		)
	}
}

func (ps PaymentStatus) String() string {
	switch ps {
	case Pending:
		return "pending"
	case Failed:
		return "failed"
	case Success:
		return "success"
	default:
		return "unknown"
	}
}

// use for database
func (ps PaymentStatus) Value() string {
	return ps.String()
}
