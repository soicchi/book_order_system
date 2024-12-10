package repository

import (
	"errors"
	"fmt"

	"event_system/internal/domain/event"
	errs "event_system/internal/errors"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventRepository struct{}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (er *EventRepository) Create(ctx echo.Context, event *event.Event) error {
	db := database.GetDB(ctx)

	eventModel := models.Event{
		ID:          event.ID(),
		Title:       event.Title(),
		Description: event.Description(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
		CreatedBy:   event.CreatedBy(),
		VenueID:     event.VenueID(),
	}

	err := db.Create(&eventModel).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.New(
			fmt.Errorf("event already exists: %s", event.ID()),
			errs.AlreadyExistError,
			errs.WithField("EventID"),
		)
	}

	if err != nil {
		return errs.New(
			fmt.Errorf("failed to create event: %s", event.ID()),
			errs.UnexpectedError,
		)
	}

	return nil
}

func (er *EventRepository) FetchByID(ctx echo.Context, eventID uuid.UUID) (*event.Event, error) {
	return nil, nil
}

func (er *EventRepository) FetchAll(ctx echo.Context) ([]*event.Event, error) {
	return nil, nil
}

func (er *EventRepository) FetchByVenueID(ctx echo.Context, venueID uuid.UUID) ([]*event.Event, error) {
	return nil, nil
}

func (er *EventRepository) Update(ctx echo.Context, event *event.Event) error {
	return nil
}
