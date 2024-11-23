package models

import (
	"time"

	"github.com/google/uuid"
)

type Shipping struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null"`         // order id
	Address   string    `gorm:"type:varchar(255);not null"` // shipping address
	Method    string    `gorm:"type:varchar(255);not null"` // shipping method
	Status    string    `gorm:"type:varchar(255);not null"` // shipping status
	CreatedAt time.Time // when the shipping was created
	UpdatedAt time.Time // when the shipping was last updated
}
