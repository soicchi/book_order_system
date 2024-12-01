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

type CancelInput struct {
	OrderID uuid.UUID
	UserID  uuid.UUID
	Details []*CancelDetailInput
}

type CancelDetailInput struct {
	BookID   uuid.UUID
	Quantity int
}

func NewCancelDetail(bookID uuid.UUID, quantity int) *CancelDetailInput {
	return &CancelDetailInput{
		BookID:   bookID,
		Quantity: quantity,
	}
}

func NewCancelInput(orderID uuid.UUID, userID uuid.UUID, details []*CancelDetailInput) *CancelInput {
	return &CancelInput{
		OrderID: orderID,
		UserID:  userID,
		Details: details,
	}
}
