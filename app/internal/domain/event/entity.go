package event

import (
	"fmt"
	"time"

	"event_system/internal/errors"

	"github.com/google/uuid"
)

type Event struct {
	id          uuid.UUID
	title       string
	description string
	startDate   time.Time
	endDate     time.Time
	createdAt   time.Time
	updatedAt   time.Time
	createdBy   uuid.UUID
	venueID     uuid.UUID
}

func new(
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	userID uuid.UUID,
	venueID uuid.UUID,
) (*Event, error) {
	// check start data
	if startDate.After(endDate) || startDate.Before(time.Now()) {
		return nil, errors.New(
			fmt.Errorf("invalid event start time: %s", startDate),
			errors.ValidationError,
			errors.WithField("StartTime"),
			errors.WithIssue(errors.InvalidTimeRange),
		)
	}

	// check end date
	if endDate.Before(startDate) || endDate.Before(time.Now()) || endDate.Equal(startDate) {
		return nil, errors.New(
			fmt.Errorf("invalid event end time: %s", endDate),
			errors.ValidationError,
			errors.WithField("EndTime"),
			errors.WithIssue(errors.InvalidTimeRange),
		)
	}

	return &Event{
		id:          uuid.New(),
		title:       title,
		description: description,
		startDate:   startDate,
		endDate:     endDate,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		createdBy:   userID,
		venueID:     venueID,
	}, nil
}

func Reconstruct(
	id uuid.UUID,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	createdBy uuid.UUID,
	venueID uuid.UUID,
) *Event {
	return &Event{
		id:          id,
		title:       title,
		description: description,
		startDate:   startDate,
		endDate:     endDate,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		createdBy:   createdBy,
		venueID:     venueID,
	}
}

func (e *Event) ID() uuid.UUID {
	return e.id
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) Description() string {
	return e.description
}

func (e *Event) StartDate() time.Time {
	return e.startDate
}

func (e *Event) EndDate() time.Time {
	return e.endDate
}

func (e *Event) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Event) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Event) CreatedBy() uuid.UUID {
	return e.createdBy
}

func (e *Event) VenueID() uuid.UUID {
	return e.venueID
}
