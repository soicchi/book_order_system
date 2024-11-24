package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID           uuid.UUID     `gorm:"type:uuid;primary_key"`
	Title        string        `gorm:"type:varchar(255);not null"`
	Author       string        `gorm:"type:varchar(255);not null"`
	Price        float64       `gorm:"type:numeric;not null"`
	CreatedAt    time.Time     `gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time     `gorm:"type:timestamp;not null"`
	OrderDetails []OrderDetail `gorm:"foreignKey:BookID"`
}
