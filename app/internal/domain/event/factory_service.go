package event

import (
	"fmt"
	"time"

	"event_system/internal/domain/user"
	"event_system/internal/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventFactory struct {
	eventRepository EventRepository
	userRepository  user.UserRepository
}

func NewEventFactory(eventRepository EventRepository, userRepository user.UserRepository) *EventFactory {
	return &EventFactory{
		eventRepository: eventRepository,
		userRepository:  userRepository,
	}
}

func (ef *EventFactory) NewEvent(
	ctx echo.Context,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	userID uuid.UUID,
	venueID uuid.UUID,
) (*Event, error) {
	// Check if the user has the right role to create an event
	user, err := ef.userRepository.FetchByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(
			fmt.Errorf("user not found: %s", userID),
			errors.NotFoundError,
			errors.WithField("UserID"),
		)
	}

	// Only organizers can create events
	if !user.IsOrganizer() {
		return nil, errors.New(
			fmt.Errorf("user is not an organizer: %s", userID),
			errors.AuthorizationError,
		)
	}

	events, err := ef.eventRepository.FetchByVenueID(ctx, venueID)
	if err != nil {
		return nil, err
	}

	// Check if the event time range conflicts with other events
	for _, event := range events {
		if event.StartDate().Before(endDate) && event.EndDate().After(startDate) {
			return nil, errors.New(
				fmt.Errorf("event time range conflict: registered: %s, new: %s", event.Title(), title),
				errors.ValidationError,
				errors.WithField("StartDateOrEndDate"),
				errors.WithIssue(errors.InvalidTimeRange),
			)
		}
	}

	event, err := new(title, description, startDate, endDate, userID, venueID)
	if err != nil {
		return nil, err
	}

	return event, nil
}
