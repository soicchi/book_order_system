package event

import (
	"fmt"
	"time"

	"event_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventUpdaterService interface {
	UpdateEvent(
		ctx echo.Context,
		eventID uuid.UUID,
		title string,
		description string,
		startDate time.Time,
		endDate time.Time,
		createdBy uuid.UUID,
		venueID uuid.UUID,
	) (*Event, error)
}

type EventUpdater struct {
	eventRepository EventRepository
}

func NewEventUpdater(
	eventRepository EventRepository,
) *EventUpdater {
	return &EventUpdater{
		eventRepository: eventRepository,
	}
}

func (eu *EventUpdater) UpdateEvent(
	ctx echo.Context,
	eventID uuid.UUID,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	createdBy uuid.UUID,
	venueID uuid.UUID,
) (*Event, error) {
	event, err := eu.eventRepository.FetchByID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New(
			fmt.Errorf("event not found: %s", eventID),
			errors.NotFoundError,
			errors.WithField("EventID"),
		)
	}

	// Check if the user has the right role to update an event
	if err := event.ValidateHost(createdBy); err != nil {
		return nil, err
	}

	events, err := eu.eventRepository.FetchByVenueID(ctx, venueID)
	if err != nil {
		return nil, err
	}

	// Update the event properties
	if err := event.SetTimeRange(startDate, endDate, events); err != nil {
		return nil, err
	}
	event.SetTitle(title)
	event.SetDescription(description)
	event.SetUpdatedAt()

	return event, nil
}
