package models

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Status       string    `gorm:"type:varchar(255);not null"`
	RegisteredAt time.Time `gorm:"type:timestamp;not null"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	EventID      uuid.UUID `gorm:"type:uuid;not null"`
	User         User
	Event        Event
	Ticket       Ticket
}
