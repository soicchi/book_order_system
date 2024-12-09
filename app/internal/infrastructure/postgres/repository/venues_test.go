package repository

import (
	"testing"
	"time"

	"event_system/internal/domain/venue"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/database/fixtures"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
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

func TestFetchAllVenues(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Fetch all venues successfully",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewVenueRepository()

			venues, repoErr := r.FetchAll(ctx)

			if tt.wantErr {
				assert.NotNil(t, repoErr)
				return
			}

			assert.Nil(t, repoErr)

			assert.Equal(t, len(fixtures.TestVenues), len(venues))
		})
	}
}

func TestFetchVenueByID(t *testing.T) {
	tests := []struct {
		name    string
		venueID uuid.UUID
		wantErr bool
	}{
		{
			name:    "Fetch venue by ID successfully",
			venueID: fixtures.TestVenues["venue1"].ID,
			wantErr: false,
		},
		{
			name:    "Fetch venue by non-existent ID",
			venueID: uuid.New(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewVenueRepository()

			venue, repoErr := r.FetchByID(ctx, tt.venueID)

			if tt.wantErr {
				assert.NotNil(t, repoErr)
				return
			}

			assert.Nil(t, repoErr)

			// Not found
			if venue == nil {
				return
			}

			assert.Equal(t, fixtures.TestVenues["venue1"].ID, venue.ID())
			assert.Equal(t, fixtures.TestVenues["venue1"].Name, venue.Name())
			assert.Equal(t, fixtures.TestVenues["venue1"].Address, venue.Address())
			assert.Equal(t, fixtures.TestVenues["venue1"].Capacity, venue.Capacity())
		})
	}
}

func TestUpdateVenue(t *testing.T) {
	tests := []struct {
		name    string
		venue   *venue.Venue
		wantErr bool
	}{
		{
			name: "Update venue successfully",
			venue: venue.Reconstruct(
				fixtures.TestVenues["venue1"].ID,
				"updated_venue",
				"updated_address",
				200,
				fixtures.TestVenues["venue1"].CreatedAt,
				time.Now(),
			),
			wantErr: false,
		},
		{
			name: "Update non-existent venue",
			venue: venue.Reconstruct(
				uuid.New(),
				"updated_venue",
				"updated_address",
				200,
				time.Now(),
				time.Now(),
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)
			db := database.GetDB(ctx)

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

			repoErr := r.Update(ctx, tt.venue)

			if tt.wantErr {
				assert.NotNil(t, repoErr)

				var venueModel models.Venue
				if err := db.Where("id = ?", tt.venue.ID()).First(&venueModel).Error; err != nil {
					return
				}

				assert.NotEqual(t, tt.venue.Name(), venueModel.Name)
				assert.NotEqual(t, tt.venue.Address(), venueModel.Address)
				assert.NotEqual(t, tt.venue.Capacity(), venueModel.Capacity)
				assert.NotEqual(t, tt.venue.UpdatedAt(), venueModel.UpdatedAt)
				return
			}

			assert.Nil(t, repoErr)

			var venueModel models.Venue
			if err := tx.Where("id = ?", tt.venue.ID()).First(&venueModel).Error; err != nil {
				t.Fatalf("failed to fetch venue: %v", err)
			}

			assert.Equal(t, tt.venue.ID(), venueModel.ID)
			assert.Equal(t, tt.venue.Name(), venueModel.Name)
			assert.Equal(t, tt.venue.Address(), venueModel.Address)
			assert.Equal(t, tt.venue.Capacity(), venueModel.Capacity)
		})
	}
}
