package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name       string    `gorm:"type:varchar(255)"`   // product name
	Price      float64   `gorm:"type:float;not null"` // product price
	CreatedAt  time.Time // when the product was created
	UpdatedAt  time.Time // when the product was last updated
	OrderItems []OrderItem
}
