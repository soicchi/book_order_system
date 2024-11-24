package users

import (
	"github.com/google/uuid"
)

type CreateInput struct {
	Username string
	Email    string
	Password string
}

func NewCreateInput(username, email, password string) *CreateInput {
	return &CreateInput{
		Username: username,
		Email:    email,
		Password: password,
	}
}

type RetrieveOutput struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewRetrieveOutput(id uuid.UUID, username, email, createdAt, updatedAt string) *RetrieveOutput {
	return &RetrieveOutput{
		ID:        id,
		Username:  username,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type ListOutput struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewListOutput(id uuid.UUID, username, email, createdAt, updatedAt string) *ListOutput {
	return &ListOutput{
		ID:        id,
		Username:  username,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type UpdateInput struct {
	ID       uuid.UUID
	Username string
	Email    string
	Password string
}

func NewUpdateInput(id uuid.UUID, username, email, password string) *UpdateInput {
	return &UpdateInput{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
}
