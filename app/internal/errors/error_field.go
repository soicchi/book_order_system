package errors

type ErrorField int

// Add ErrorField when you need a new error field
const (
	NoField ErrorField = iota
	// User
	User
	Username
	UserEmail
	UserPassword
	// Book
	Book
	BookTitle
	BookAuthor
	BookPrice
	BookStock
	// OrderDetail
	OrderDetail
	OrderDetailQuantity
	OrderDetailPrice
	// Order
	Order
	OrderTotalPrice
	OrderStatus
)

func (e ErrorField) String() string {
	switch e {
	case User:
		return "User"
	case Username:
		return "Username"
	case UserEmail:
		return "UserEmail"
	case UserPassword:
		return "UserPassword"
	case Book:
		return "Book"
	case BookTitle:
		return "BookTitle"
	case BookAuthor:
		return "BookAuthor"
	case BookPrice:
		return "BookPrice"
	case BookStock:
		return "BookStock"
	case OrderDetailQuantity:
		return "OrderDetailQuantity"
	case OrderDetailPrice:
		return "OrderDetailPrice"
	case Order:
		return "Order"
	case OrderTotalPrice:
		return "OrderTotalPrice"
	case OrderStatus:
		return "OrderStatus"
	default:
		return "Unknown"
	}
}
