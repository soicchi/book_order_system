package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	StartDate   time.Time `gorm:"type:timestamp;not null"`
	EndDate     time.Time `gorm:"type:timestamp;not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null"`
	VenueID     uuid.UUID `gorm:"type:uuid;not null"`
	Venue       Venue
}
