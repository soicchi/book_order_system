package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`        // customer name
	Email     string    `gorm:"type:varchar(255);unique"` // customer email
	CreatedAt time.Time // when the customer was created
	UpdatedAt time.Time // when the customer was last updated
}
