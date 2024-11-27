package orders

import (
	"github.com/google/uuid"
)

type CreateRequest struct {
	UserID       uuid.UUID              `param:"user_id" validate:"required,uuid"`
	OrderID      uuid.UUID              `param:"order_id" validate:"required,uuid"`
	OrderDetails []*CreateDetailRequest `json:"order_details" validate:"required,dive"`
}

type CreateDetailRequest struct {
	BookID   uuid.UUID `json:"book_id" validate:"required,uuid"`
	Quantity int       `json:"quantity" validate:"required"`
	Price    float64   `json:"price" validate:"required"`
}

type CancelRequest struct {
	OrderID uuid.UUID `param:"order_id" validate:"required,uuid"`
}
