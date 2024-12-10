package repository

import (
	"testing"
	"time"

	"event_system/internal/domain/event"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/database/fixtures"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	tests := []struct {
		name    string
		event   *event.Event
		wantErr bool
	}{
		{
			name: "Create event successfully",
			event: event.Reconstruct(
				uuid.New(),
				"test_event",
				"test_description",
				time.Date(2024, time.November, 12, 10, 0, 0, 0, time.UTC),
				time.Date(2024, time.November, 12, 10, 23, 59, 0, time.UTC),
				time.Now(),
				time.Now(),
				fixtures.TestUsers["organizer1"].ID,
				fixtures.TestVenues["venue1"].ID,
			),
			wantErr: false,
		},
		{
			name: "Create event with duplicated ID",
			event: event.Reconstruct(
				fixtures.TestEvents["event1"].ID, // Duplicated ID
				"test_event",
				"test_description",
				time.Date(2024, time.November, 12, 10, 0, 0, 0, time.UTC),
				time.Date(2024, time.November, 12, 10, 23, 59, 0, time.UTC),
				time.Now(),
				time.Now(),
				fixtures.TestUsers["organizer1"].ID,
				fixtures.TestVenues["venue1"].ID,
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)
			db := database.GetDB(ctx)

			// Fetch events before test
			var beforeEventModels []models.Event
			db.Find(&beforeEventModels)

			// Start transaction
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("Failed to begin transaction: %v", tx.Error)
			}

			defer func() {
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("Failed to rollback transaction: %v", err)
				}
			}()

			r := NewEventRepository()

			repoErr := r.Create(ctx, tt.event)

			if tt.wantErr {
				assert.NotNil(t, repoErr)

				var afterEventModels []models.Event
				if err := db.Find(&afterEventModels).Error; err != nil {
					t.Fatalf("Failed to fetch events: %v", err)
				}

				assert.Equal(t, len(afterEventModels), len(beforeEventModels))
				return
			}

			assert.Nil(t, repoErr)

			var eventModel models.Event
			if err := tx.First(&eventModel, "id = ?", tt.event.ID()).Error; err != nil {
				t.Fatalf("Failed to fetch event: %v", err)
			}

			assert.Equal(t, tt.event.ID(), eventModel.ID)
			assert.Equal(t, tt.event.Title(), eventModel.Title)
			assert.Equal(t, tt.event.Description(), eventModel.Description)
			assert.Equal(t, tt.event.StartDate().Unix(), eventModel.StartDate.Unix())
			assert.Equal(t, tt.event.EndDate().Unix(), eventModel.EndDate.Unix())
			assert.Equal(t, tt.event.CreatedBy(), eventModel.CreatedBy)
			assert.Equal(t, tt.event.VenueID(), eventModel.VenueID)
		})
	}
}

func TestFetchEventByID(t *testing.T) {
	tests := []struct {
		name    string
		eventID uuid.UUID
		wantErr bool
	}{
		{
			name:    "Fetch event by ID successfully",
			eventID: fixtures.TestEvents["event1"].ID,
			wantErr: false,
		},
		{
			name:    "Fetch event by non-existent ID",
			eventID: uuid.New(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewEventRepository()

			event, repoErr := r.FetchByID(ctx, tt.eventID)

			if tt.wantErr {
				assert.Nil(t, event)
				assert.NotNil(t, repoErr)
				return
			}

			// The event is not found
			if event == nil {
				return
			}

			assert.NotNil(t, event)
			assert.Nil(t, repoErr)

			assert.Equal(t, tt.eventID, event.ID())
			assert.Equal(t, fixtures.TestEvents["event1"].Title, event.Title())
			assert.Equal(t, fixtures.TestEvents["event1"].Description, event.Description())
			assert.Equal(t, fixtures.TestEvents["event1"].StartDate.Unix(), event.StartDate().Unix())
			assert.Equal(t, fixtures.TestEvents["event1"].EndDate.Unix(), event.EndDate().Unix())
			assert.Equal(t, fixtures.TestEvents["event1"].CreatedBy, event.CreatedBy())
			assert.Equal(t, fixtures.TestEvents["event1"].VenueID, event.VenueID())
		})
	}
}

func TestFetchAllEvents(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Fetch all events successfully",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewEventRepository()

			events, repoErr := r.FetchAll(ctx)

			if tt.wantErr {
				assert.Nil(t, events)
				assert.NotNil(t, repoErr)
				return
			}

			assert.NotNil(t, events)
			assert.Nil(t, repoErr)

			assert.Equal(t, len(fixtures.TestEvents), len(events))
		})
	}
}
