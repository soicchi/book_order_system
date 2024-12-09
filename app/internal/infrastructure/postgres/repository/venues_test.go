package repository

import (
	"testing"
	"time"

	"event_system/internal/domain/venue"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/database/fixtures"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateVenue(t *testing.T) {
	tests := []struct {
		name    string
		venue   *venue.Venue
		wantErr bool
	}{
		{
			name: "Create venue successfully",
			venue: venue.New(
				"test_venue",
				"test_address",
				100,
			),
			wantErr: false,
		},
		{
			name: "Create venue with duplicated ID",
			venue: venue.Reconstruct(
				fixtures.TestVenues["venue1"].ID,
				"test_venue",
				"test_address",
				100,
				time.Now(),
				time.Now(),
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		e := echo.New()
		ctx := e.NewContext(nil, nil)
		db := database.GetDB(ctx)

		// Fetch venues before test
		var beforeVenueModels []models.Venue
		if err := db.Find(&beforeVenueModels).Error; err != nil {
			t.Fatalf("failed to fetch venues: %v", err)
		}

		// Start transaction
		tx, err := database.BeginTx(ctx)
		if err != nil {
			t.Fatalf("Failed to begin transaction: %v", tx.Error)
		}

		// Rollback transaction at the end of the test
		defer func() {
			if err := tx.Rollback().Error; err != nil {
				t.Fatalf("failed to rollback transaction: %v", err)
			}
		}()

		r := NewVenueRepository()

		repoErr := r.Create(ctx, tt.venue)

		if tt.wantErr {
			assert.NotNil(t, repoErr)

			var afterVenueModels []models.Venue
			if err := db.Find(&afterVenueModels).Error; err != nil {
				t.Fatalf("failed to fetch venues: %v", err)
			}

			assert.Equal(t, len(beforeVenueModels), len(afterVenueModels))
			return
		}

		assert.Nil(t, repoErr)

		var venueModels models.Venue
		if err := tx.Where("id = ?", tt.venue.ID()).First(&venueModels).Error; err != nil {
			t.Fatalf("failed to fetch venue: %v", err)
		}

		assert.Equal(t, tt.venue.ID(), venueModels.ID)
		assert.Equal(t, tt.venue.Name(), venueModels.Name)
		assert.Equal(t, tt.venue.Address(), venueModels.Address)
		assert.Equal(t, tt.venue.Capacity(), venueModels.Capacity)
	}
}
