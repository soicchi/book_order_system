package models

import (
	"github.com/google/uuid"
)

type OrderDetail struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	OrderID  uuid.UUID `gorm:"type:uuid;not null"`
	BookID   uuid.UUID `gorm:"type:uuid;not null"`
	Quantity int       `gorm:"type:integer;not null"`
	Price    float64   `gorm:"type:numeric;not null"`
	Book     Book
	Order    Order
}
