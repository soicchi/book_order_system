package values

import (
	"fmt"

	"github.com/soicchi/book_order_system/internal/errors"
)

type PaymentMethod int

const (
	CreditCard PaymentMethod = iota + 1
	BankTransfer
	ConvenienceStore
)

func NewPaymentMethod(method string) (PaymentMethod, error) {
	switch method {
	case "credit_card":
		return CreditCard, nil
	case "bank_transfer":
		return BankTransfer, nil
	case "convenience_store":
		return ConvenienceStore, nil
	default:
		return 0, errors.New(
			fmt.Errorf("invalid payment method: %s", method),
			errors.InvalidRequest,
		)
	}
}

func (pm PaymentMethod) String() string {
	switch pm {
	case CreditCard:
		return "credit_card"
	case BankTransfer:
		return "bank_transfer"
	case ConvenienceStore:
		return "convenience_store"
	default:
		return "unknown"
	}
}
