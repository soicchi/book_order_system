package models

import (
	"time"

	"github.com/google/uuid"
)

type Venue struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Address   string    `gorm:"type:text;not null"`
	Capacity  int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	Events    []Event   `gorm:"foreignKey:VenueID"`
}
