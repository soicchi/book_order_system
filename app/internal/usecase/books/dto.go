package books

import (
	"github.com/google/uuid"
)

type CreateInput struct {
	Title  string
	Author string
	Price  float64
	Stock  int
}

func NewCreateInput(title, author string, price float64, stock int) *CreateInput {
	return &CreateInput{
		Title:  title,
		Author: author,
		Price:  price,
		Stock:  stock,
	}
}

type RetrieveOutput struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewRetrieveOutput(
	id uuid.UUID,
	title string,
	author string,
	price float64,
	stock int,
	createdAt string,
	updatedAt string,
) *RetrieveOutput {
	return &RetrieveOutput{
		ID:        id,
		Title:     title,
		Author:    author,
		Price:     price,
		Stock:     stock,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type ListOutput struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewListOutput(
	id uuid.UUID,
	title string,
	author string,
	price float64,
	stock int,
	createdAt string,
	updatedAt string,
) *ListOutput {
	return &ListOutput{
		ID:        id,
		Title:     title,
		Author:    author,
		Price:     price,
		Stock:     stock,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type UpdateInput struct {
	ID     uuid.UUID
	Title  string
	Author string
	Price  float64
	Stock  int
}

func NewUpdateInput(id uuid.UUID, title, author string, price float64, stock int) *UpdateInput {
	return &UpdateInput{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
		Stock:  stock,
	}
}
