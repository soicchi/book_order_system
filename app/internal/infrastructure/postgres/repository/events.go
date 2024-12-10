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
	db := database.GetDB(ctx)

	var eventModel models.Event
	err := db.Where("id = ?", eventID).First(&eventModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch event: %s", eventID),
			errs.UnexpectedError,
		)
	}

	return event.Reconstruct(
		eventModel.ID,
		eventModel.Title,
		eventModel.Description,
		eventModel.StartDate,
		eventModel.EndDate,
		eventModel.CreatedAt,
		eventModel.UpdatedAt,
		eventModel.CreatedBy,
		eventModel.VenueID,
	), nil
}

func (er *EventRepository) FetchAll(ctx echo.Context) ([]*event.Event, error) {
	db := database.GetDB(ctx)

	var eventModels []models.Event
	if err := db.Find(&eventModels).Error; err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch events"),
			errs.UnexpectedError,
		)
	}

	events := make([]*event.Event, 0, len(eventModels))
	for _, eventModel := range eventModels {
		events = append(events, event.Reconstruct(
			eventModel.ID,
			eventModel.Title,
			eventModel.Description,
			eventModel.StartDate,
			eventModel.EndDate,
			eventModel.CreatedAt,
			eventModel.UpdatedAt,
			eventModel.CreatedBy,
			eventModel.VenueID,
		))
	}

	return events, nil
}

func (er *EventRepository) FetchByVenueID(ctx echo.Context, venueID uuid.UUID) ([]*event.Event, error) {
	db := database.GetDB(ctx)

	var eventModels []models.Event
	err := db.Where("venue_id = ?", venueID).Find(&eventModels).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch events by venue: %s", venueID),
			errs.UnexpectedError,
		)
	}

	events := make([]*event.Event, 0, len(eventModels))
	for _, eventModel := range eventModels {
		events = append(events, event.Reconstruct(
			eventModel.ID,
			eventModel.Title,
			eventModel.Description,
			eventModel.StartDate,
			eventModel.EndDate,
			eventModel.CreatedAt,
			eventModel.UpdatedAt,
			eventModel.CreatedBy,
			eventModel.VenueID,
		))
	}

	return events, nil
}

func (er *EventRepository) Update(ctx echo.Context, event *event.Event) error {
	db := database.GetDB(ctx)

	var eventModel models.Event
	result := db.Model(&eventModel).Where("id = ?", event.ID()).Updates(models.Event{
		Title:       event.Title(),
		Description: event.Description(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
		UpdatedAt:   event.UpdatedAt(),
	})
	if result.Error != nil {
		return errs.New(
			fmt.Errorf("failed to update event: %s", event.ID()),
			errs.UnexpectedError,
		)
	}

	if result.RowsAffected == 0 {
		return errs.New(
			fmt.Errorf("event not found: %s", event.ID()),
			errs.NotFoundError,
			errs.WithField("EventID"),
		)
	}

	return nil
}
