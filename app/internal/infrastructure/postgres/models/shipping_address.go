package models

import (
	"time"

	"github.com/google/uuid"
)

type ShippingAddress struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
	Prefecture string    `gorm:"type:varchar(255);not null"`
	City       string    `gorm:"type:varchar(255);not null"`
	State      string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;not null"`
	Customer   Customer
}
