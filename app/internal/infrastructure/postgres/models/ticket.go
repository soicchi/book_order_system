package models

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key"`
	QRCode         string    `gorm:"type:varchar(255);not null"`
	Status         string    `gorm:"type:varchar(255);not null"`
	IssuedAt       time.Time `gorm:"type:timestamp;not null"`
	UsedAt         time.Time `gorm:"type:timestamp"`
	RegistrationID uuid.UUID `gorm:"type:uuid;not null"`
}
