package models

import (
	"github.com/google/uuid"
)

type Customer struct {
	CustomerID uuid.UUID `gorm:"type:uuid;primaryKey;default"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Email      string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	TimeStamps `gorm:"embedded"`
}
