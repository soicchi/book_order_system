package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null"` // order id
	ProductID uuid.UUID `gorm:"type:uuid;not null"` // product id
	Quantity  int       `gorm:"type:int;not null"`  // quantity of product
	Order     Order
	Product   Product
}
