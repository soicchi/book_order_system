package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Status     string    `gorm:"type:varchar(255);not null"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
	TimeStamp  `gorm:"embedded"`
}
