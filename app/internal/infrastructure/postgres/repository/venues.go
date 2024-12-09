package repository

import (
	"errors"
	"fmt"

	"event_system/internal/domain/venue"
	errs "event_system/internal/errors"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type VenueRepository struct{}

func NewVenueRepository() *VenueRepository {
	return &VenueRepository{}
}

func (r *VenueRepository) Create(ctx echo.Context, venue *venue.Venue) error {
	db := database.GetDB(ctx)

	venueModel := &models.Venue{
		ID:        venue.ID(),
		Name:      venue.Name(),
		Address:   venue.Address(),
		Capacity:  venue.Capacity(),
		CreatedAt: venue.CreatedAt(),
		UpdatedAt: venue.UpdatedAt(),
	}

	err := db.Create(venueModel).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.New(
			fmt.Errorf("venue id already exists: %w", err),
			errs.AlreadyExistError,
			errs.WithField("ID"),
		)
	}

	if err != nil {
		return errs.New(
			fmt.Errorf("failed to create venue: %w", err),
			errs.UnexpectedError,
		)
	}

	return nil
}

func (r *VenueRepository) FetchAll(ctx echo.Context) ([]*venue.Venue, error) {
	db := database.GetDB(ctx)

	var venueModels []*models.Venue
	err := db.Find(&venueModels).Error
	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch venues: %w", err),
			errs.UnexpectedError,
		)
	}

	var venues []*venue.Venue
	for _, venueModel := range venueModels {
		venues = append(venues, venue.Reconstruct(
			venueModel.ID,
			venueModel.Name,
			venueModel.Address,
			venueModel.Capacity,
			venueModel.CreatedAt,
			venueModel.UpdatedAt,
		))
	}

	return venues, nil
}

func (r *VenueRepository) FetchByID(ctx echo.Context, venueID uuid.UUID) (*venue.Venue, error) {
	db := database.GetDB(ctx)

	var venueModel models.Venue
	err := db.Where("id = ?", venueID).First(&venueModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch venue: %w", err),
			errs.UnexpectedError,
		)
	}

	return venue.Reconstruct(
		venueModel.ID,
		venueModel.Name,
		venueModel.Address,
		venueModel.Capacity,
		venueModel.CreatedAt,
		venueModel.UpdatedAt,
	), nil
}

func (r *VenueRepository) Update(ctx echo.Context, venue *venue.Venue) error {
	db := database.GetDB(ctx)

	var venueModel models.Venue
	result := db.Model(&venueModel).Where("id = ?", venue.ID()).Updates(models.Venue{
		Name:      venue.Name(),
		Address:   venue.Address(),
		Capacity:  venue.Capacity(),
		UpdatedAt: venue.UpdatedAt(),
	})
	if result.Error != nil {
		return errs.New(
			fmt.Errorf("failed to update venue: %w", result.Error),
			errs.UnexpectedError,
		)
	}

	if result.RowsAffected == 0 {
		return errs.New(
			fmt.Errorf("venue not found: %w", result.Error),
			errs.NotFoundError,
			errs.WithField("VenueID"),
		)
	}

	return nil
}
