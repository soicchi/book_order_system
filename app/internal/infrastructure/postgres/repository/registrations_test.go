package repository

import (
	"testing"
	"time"

	"event_system/internal/domain/registration"
	"event_system/internal/domain/status"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/database/fixtures"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateRegistration(t *testing.T) {
	tests := []struct {
		name         string
		registration *registration.Registration
		wantErr      bool
	}{
		{
			name: "Create registration successfully",
			registration: registration.New(
				fixtures.TestEvents["event1"].ID,
				fixtures.TestUsers["attendee1"].ID,
			),
			wantErr: false,
		},
		{
			name: "Create registration with duplicated ID",
			registration: registration.Reconstruct(
				fixtures.TestRegistrations["registration1"].ID,
				time.Now(),
				status.New(status.Registered),
				fixtures.TestRegistrations["registration1"].UserID,
				fixtures.TestRegistrations["registration1"].EventID,
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)
			db := database.GetDB(ctx)

			var beforeRegistrationModels []models.Registration
			if err := db.Find(&beforeRegistrationModels).Error; err != nil {
				t.Fatalf("failed to fetch registrations: %v", err)
			}

			// Start transaction
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("failed to begin transaction: %v", err)
			}

			// Rollback transaction at the end of the test
			defer func() {
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("failed to rollback transaction: %v", err)
				}
			}()

			r := NewRegistrationRepository()

			repoErr := r.Create(ctx, tt.registration)

			if tt.wantErr {
				assert.NotNil(t, repoErr)

				var afterRegistrationModels []models.Registration
				if err = db.Find(&afterRegistrationModels).Error; err != nil {
					t.Fatalf("failed to fetch registrations: %v", err)
				}

				assert.Equal(t, len(beforeRegistrationModels), len(afterRegistrationModels))
				return
			}

			assert.Nil(t, repoErr)

			var registrationModel models.Registration
			if err := tx.Where("id = ?", tt.registration.ID()).First(&registrationModel).Error; err != nil {
				t.Fatalf("failed to fetch registration: %v", err)
			}

			assert.Equal(t, tt.registration.ID(), registrationModel.ID)
			assert.Equal(t, tt.registration.Status().Value().String(), registrationModel.Status)
			assert.Equal(t, tt.registration.RegisteredAt().Unix(), registrationModel.RegisteredAt.Unix())
			assert.Equal(t, tt.registration.UserID(), registrationModel.UserID)
			assert.Equal(t, tt.registration.EventID(), registrationModel.EventID)
		})
	}
}

func TestFetchRegistrationsByEventID(t *testing.T) {
	tests := []struct {
		name    string
		eventID uuid.UUID
	}{
		{
			name:    "Fetch registrations by event ID successfully",
			eventID: fixtures.TestEvents["event1"].ID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewRegistrationRepository()

			registrations, err := r.FetchByEventID(ctx, tt.eventID)

			assert.Nil(t, err)
			assert.NotEmpty(t, registrations)
		})
	}
}

func TestUpdateRegistration(t *testing.T) {
	tests := []struct {
		name         string
		registration *registration.Registration
		wantErr      bool
	}{
		{
			name: "Update registration successfully",
			registration: registration.Reconstruct(
				fixtures.TestRegistrations["registration1"].ID,
				fixtures.TestRegistrations["registration1"].RegisteredAt,
				status.New(status.Canceled),
				fixtures.TestRegistrations["registration1"].UserID,
				fixtures.TestRegistrations["registration1"].EventID,
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)
			db := database.GetDB(ctx)

			var beforeRegistrationModel models.Registration
			db.Where("id = ?", tt.registration.ID()).First(&beforeRegistrationModel)

			// Start transaction
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("failed to begin transaction: %v", err)
			}

			// Rollback transaction at the end of the test
			defer func() {
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("failed to rollback transaction: %v", err)
				}
			}()

			r := NewRegistrationRepository()

			repoErr := r.Update(ctx, tt.registration)

			if tt.wantErr {
				assert.NotNil(t, repoErr)

				var afterRegistrationModel models.Registration
				if err = db.Preload("Ticket").Where("id = ?", tt.registration.ID()).First(&afterRegistrationModel).Error; err != nil {
					t.Fatalf("failed to fetch registrations: %v", err)
				}

				assert.NotEqual(t, tt.registration.Status().Value().String(), afterRegistrationModel.Status)

				return
			}

			assert.Nil(t, repoErr)

			var registrationModel models.Registration
			if err := tx.Where("id = ?", tt.registration.ID()).First(&registrationModel).Error; err != nil {
				t.Fatalf("failed to fetch registration: %v", err)
			}

			assert.Equal(t, tt.registration.Status().Value().String(), registrationModel.Status)
		})
	}
}
