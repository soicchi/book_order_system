package models

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default"`
	Name              string    `gorm:"type:varchar(255);not null"`
	Email             string    `gorm:"type:varchar(255);not null;unique"`
	Password          string    `gorm:"type:varchar(255);not null"`
	TimeStamp         `gorm:"embedded"`
	ShippingAddresses []ShippingAddress `gorm:"constraint:OnDelete:CASCADE"`
}
