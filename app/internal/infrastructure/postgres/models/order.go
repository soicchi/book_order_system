package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	TotalPrice   float64   `gorm:"type:numeric;not null"`
	OrderedAt    time.Time `gorm:"type:timestamp;not null"`
	User         User
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}
