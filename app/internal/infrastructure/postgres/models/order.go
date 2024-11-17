package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID        uuid.UUID `gorm:"type:uuid;not null;index"`
	ShippingAddressID uuid.UUID `gorm:"type:uuid;not null;index"`
	TimeStamp         `gorm:"embedded"`
}
