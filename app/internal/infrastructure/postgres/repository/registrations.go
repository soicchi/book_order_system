package repository

import (
	"fmt"

	"event_system/internal/domain/registration"
	"event_system/internal/domain/status"
	errs "event_system/internal/errors"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RegistrationRepository struct{}

func NewRegistrationRepository() *RegistrationRepository {
	return &RegistrationRepository{}
}

func (r *RegistrationRepository) Create(ctx echo.Context, registration *registration.Registration) error {
	db := database.GetDB(ctx)

	registrationModel := models.Registration{
		ID:           registration.ID(),
		Status:       registration.Status().Value().String(),
		RegisteredAt: registration.RegisteredAt(),
		UserID:       registration.UserID(),
		EventID:      registration.EventID(),
	}

	err := db.Create(&registrationModel).Error
	if err != nil {
		return errs.New(
			fmt.Errorf("failed to create registration: %w", err),
			errs.UnexpectedError,
		)
	}

	return nil
}

func (r *RegistrationRepository) FetchByEventID(ctx echo.Context, eventID uuid.UUID) ([]*registration.Registration, error) {
	db := database.GetDB(ctx)

	var registrationModels []*models.Registration
	err := db.Preload("Ticket").Where("event_id = ?", eventID).Find(&registrationModels).Error
	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch registrations: %w", err),
			errs.UnexpectedError,
		)
	}

	var registrations []*registration.Registration
	for _, registrationModel := range registrationModels {
		registrations = append(registrations, registration.Reconstruct(
			registrationModel.ID,
			registrationModel.RegisteredAt,
			status.Reconstruct(status.FromString(registrationModel.Status)),
			registrationModel.UserID,
			registrationModel.EventID,
		))
	}

	return registrations, nil
}

func (r *RegistrationRepository) Update(ctx echo.Context, registration *registration.Registration) error {
	db := database.GetDB(ctx)

	result := db.Model(&models.Registration{}).Where("id = ?", registration.ID()).Updates(models.Registration{
		Status: registration.Status().Value().String(),
	})
	if result.Error != nil {
		return errs.New(
			fmt.Errorf("failed to update registration: %w", result.Error),
			errs.UnexpectedError,
		)
	}

	if result.RowsAffected == 0 {
		return errs.New(
			fmt.Errorf("registration %s not found", registration.ID()),
			errs.NotFoundError,
			errs.WithField("RegistrationID"),
		)
	}

	return nil
}
