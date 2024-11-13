package models

import (
	"github.com/google/uuid"
)

type ShippingAddress struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
	Prefecture string    `gorm:"type:varchar(255);not null"`
	City       string    `gorm:"type:varchar(255);not null"`
	State      string    `gorm:"type:varchar(255);not null"`
	TimeStamp  `gorm:"embedded"`
	Customer   Customer
}
