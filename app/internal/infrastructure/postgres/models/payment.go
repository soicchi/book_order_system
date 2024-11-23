package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null"`  // order id
	Amount    float64   `gorm:"type:float;not null"` // payment amount
	CreatedAt time.Time // when the payment was created
	UpdatedAt time.Time // when the payment was last updated
}
