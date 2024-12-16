package event

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	targetDate := time.Now().AddDate(0, 0, 1)

	tests := []struct {
		name        string
		title       string
		description string
		startDate   time.Time
		endDate     time.Time
		userID      uuid.UUID
		venueID     uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Create event entity successfully",
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			userID:      uuid.New(),
			venueID:     uuid.New(),
		},
		{
			name:        "Create event entity failed with start date after end date",
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 2),
			userID:      uuid.New(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name:        "Create event entity failed with start date before current time",
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, -1),
			endDate:     time.Now().AddDate(0, 0, 2),
			userID:      uuid.New(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name:        "Create event entity failed with end date before start date",
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 2),
			userID:      uuid.New(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name:        "Create event entity failed with end date before current time",
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, -1),
			userID:      uuid.New(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name:        "Create event entity failed with end date equal start date",
			title:       "test_title",
			description: "test_description",
			startDate:   targetDate,
			endDate:     targetDate,
			userID:      uuid.New(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			event, err := new(tt.title, tt.description, tt.startDate, tt.endDate, tt.userID, tt.venueID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, event)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, event)
			assert.Equal(t, tt.title, event.Title())
			assert.Equal(t, tt.description, event.Description())
			assert.Equal(t, tt.startDate.Unix(), event.StartDate().Unix())
			assert.Equal(t, tt.endDate.Unix(), event.EndDate().Unix())
			assert.Equal(t, tt.userID, event.CreatedBy())
			assert.Equal(t, tt.venueID, event.VenueID())
		})
	}
}

func TestSetTimeRange(t *testing.T) {
	events := []*Event{
		Reconstruct(
			uuid.New(),
			"test_title",
			"test_description",
			time.Now().AddDate(0, 0, 1),
			time.Now().AddDate(0, 0, 2),
			time.Now(),
			time.Now(),
			uuid.New(),
			uuid.New(),
		),
		Reconstruct(
			uuid.New(),
			"test_title",
			"test_description",
			time.Now().AddDate(0, 0, 5),
			time.Now().AddDate(0, 0, 6),
			time.Now(),
			time.Now(),
			uuid.New(),
			uuid.New(),
		),
	}

	tests := []struct {
		name      string
		startDate time.Time
		endDate   time.Time
		event     *Event
		wantErr   bool
	}{
		{
			name:      "Set time range successfully",
			startDate: time.Now().AddDate(0, 0, 3),
			endDate:   time.Now().AddDate(0, 0, 4),
			event: Reconstruct(
				uuid.New(),
				"test_title",
				"test_description",
				time.Now().AddDate(0, 0, 1),
				time.Now().AddDate(0, 0, 2),
				time.Now(),
				time.Now(),
				uuid.New(),
				uuid.New(),
			),
			wantErr: false,
		},
		{
			name:      "Set time range failed with start date after end date",
			startDate: time.Now().AddDate(0, 0, 4),
			endDate:   time.Now().AddDate(0, 0, 3),
			event: Reconstruct(
				uuid.New(),
				"test_title",
				"test_description",
				time.Now().AddDate(0, 0, 1),
				time.Now().AddDate(0, 0, 2),
				time.Now(),
				time.Now(),
				uuid.New(),
				uuid.New(),
			),
			wantErr: true,
		},
		{
			name:      "Set time range failed with start date before current time",
			startDate: time.Now().AddDate(0, 0, -1),
			endDate:   time.Now().AddDate(0, 0, 3),
			event: Reconstruct(
				uuid.New(),
				"test_title",
				"test_description",
				time.Now().AddDate(0, 0, 1),
				time.Now().AddDate(0, 0, 2),
				time.Now(),
				time.Now(),
				uuid.New(),
				uuid.New(),
			),
			wantErr: true,
		},
		{
			name:      "Set time range failed with invalid time range",
			startDate: time.Now().AddDate(0, 0, 4),
			endDate:   time.Now().AddDate(0, 0, 6),
			event: Reconstruct(
				uuid.New(),
				"test_title",
				"test_description",
				time.Now().AddDate(0, 0, 1),
				time.Now().AddDate(0, 0, 2),
				time.Now(),
				time.Now(),
				uuid.New(),
				uuid.New(),
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.event.SetTimeRange(tt.startDate, tt.endDate, events)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.startDate, tt.event.StartDate())
			assert.Equal(t, tt.endDate, tt.event.EndDate())
		})
	}
}

func TestValidateHost(t *testing.T) {
	userID := uuid.New()

	tests := []struct {
		name    string
		userID  uuid.UUID
		event   *Event
		wantErr bool
	}{
		{
			name:   "Validate host successfully",
			userID: userID,
			event: Reconstruct(
				uuid.New(),
				"test_title",
				"test_description",
				time.Now().AddDate(0, 0, 1),
				time.Now().AddDate(0, 0, 2),
				time.Now(),
				time.Now(),
				userID,
				uuid.New(),
			),
			wantErr: false,
		},
		{
			name:   "Validate host failed with invalid user",
			userID: uuid.New(),
			event: Reconstruct(
				uuid.New(),
				"test_title",
				"test_description",
				time.Now().AddDate(0, 0, 1),
				time.Now().AddDate(0, 0, 2),
				time.Now(),
				time.Now(),
				uuid.New(),
				uuid.New(),
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.event.ValidateHost(tt.userID)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
