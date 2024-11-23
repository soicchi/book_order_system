package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID  uuid.UUID `gorm:"type:uuid;not null"`  // customer id
	TotalAmount float64   `gorm:"type:float;not null"` // total amount of order
	Customer    Customer
	Payment     Payment
	Shipping    Shipping
}
