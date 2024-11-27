package books

import (
	"github.com/google/uuid"
)

type CreateRequest struct {
	Title  string  `json:"title" validate:"required"`
	Author string  `json:"author" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Stock  int     `json:"stock" validate:"required"`
}

type RetrieveRequest struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

type UpdateRequest struct {
	ID     uuid.UUID `param:"id" validate:"required,uuid"`
	Title  string    `json:"title" validate:"required"`
	Author string    `json:"author" validate:"required"`
	Price  float64   `json:"price" validate:"required"`
	Stock  int       `json:"stock" validate:"required"`
}

func NewUpdateRequest(id uuid.UUID, title, author string, price float64, stock int) *UpdateRequest {
	return &UpdateRequest{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
		Stock:  stock,
	}
}
