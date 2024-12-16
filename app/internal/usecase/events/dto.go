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

type UpdateInput struct {
	EventID     uuid.UUID
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	CreatedBy   uuid.UUID
	VenueID     uuid.UUID
}

func NewUpdateInput(eventID uuid.UUID, title, description string, startDate, endDate time.Time, createdBy, venueID uuid.UUID) *UpdateInput {
	return &UpdateInput{
		EventID:     eventID,
		Title:       title,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedBy:   createdBy,
		VenueID:     venueID,
	}
}
