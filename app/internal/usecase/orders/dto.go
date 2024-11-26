package orders

import (
	"github.com/google/uuid"
)

type CreateDetailInput struct {
	BookID   uuid.UUID
	Quantity int
	Price    float64
}

type CreateInput struct {
	UserID       uuid.UUID
	OrderDetails []*CreateDetailInput
}

func NewCreateInput(userID uuid.UUID, orderDetails []*CreateDetailInput) *CreateInput {
	return &CreateInput{
		UserID:       userID,
		OrderDetails: orderDetails,
	}
}

func NewCreateDetailInput(bookID uuid.UUID, quantity int, price float64) *CreateDetailInput {
	return &CreateDetailInput{
		BookID:   bookID,
		Quantity: quantity,
		Price:    price,
	}
}
