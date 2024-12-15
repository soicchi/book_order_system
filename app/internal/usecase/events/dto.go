package events

import (
	"time"

	"github.com/google/uuid"
)

type CreateInput struct {
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	CreatedBy   uuid.UUID
	VenueID     uuid.UUID
}

func NewCreateInput(title, description string, startDate, endDate time.Time, createdBy, venueID uuid.UUID) *CreateInput {
	return &CreateInput{
		Title:       title,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedBy:   createdBy,
		VenueID:     venueID,
	}
}
