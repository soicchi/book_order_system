package dto

import (
	"time"
)

type CreateCustomerInput struct {
	Name     string
	Email    string
	Password string
}

type FetchCustomerOutput struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
