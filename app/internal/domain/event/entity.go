package event

import (
	"fmt"
	"time"

	"event_system/internal/domain/user"
	"event_system/internal/errors"

	"github.com/google/uuid"
)

type Event struct {
	id          uuid.UUID
	title       string
	description string
	startTime   time.Time
	endTime     time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func New(
	title string,
	description string,
	startTime time.Time,
	endTime time.Time,
) (*Event, error) {
	if startTime.After(endTime) || startTime.Before(time.Now()) {
		return nil, errors.New(
			fmt.Errorf("invalid event start time: %s", startTime),
			errors.ValidationError,
			errors.WithField("StartTime"),
			errors.WithIssue(errors.InvalidTimeRange),
		)
	}

	if endTime.Before(startTime) || endTime.Before(time.Now()) || endTime.Equal(startTime) {
		return nil, errors.New(
			fmt.Errorf("invalid event end time: %s", endTime),
			errors.ValidationError,
			errors.WithField("EndTime"),
			errors.WithIssue(errors.InvalidTimeRange),
		)
	}

	// Validate createdBy

	return &Event{
		id:          uuid.New(),
		title:       title,
		description: description,
		startTime:   startTime,
		endTime:     endTime,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func Reconstruct(
	id uuid.UUID,
	title string,
	description string,
	startTime time.Time,
	endTime time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	createdBy *user.User,
) *Event {
	return &Event{
		id:          id,
		title:       title,
		description: description,
		startTime:   startTime,
		endTime:     endTime,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
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

func (e *Event) StartTime() time.Time {
	return e.startTime
}

func (e *Event) EndTime() time.Time {
	return e.endTime
}

func (e *Event) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Event) UpdatedAt() time.Time {
	return e.updatedAt
}
