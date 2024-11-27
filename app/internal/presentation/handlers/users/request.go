package users

import (
	"github.com/google/uuid"
)

type CreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RetrieveRequest struct {
	ID uuid.UUID `param:"user_id" validate:"required,uuid"`
}

type UpdateRequest struct {
	ID       uuid.UUID `param:"id" validate:"required,uuid"`
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required"`
}

type DeleteRequest struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}
