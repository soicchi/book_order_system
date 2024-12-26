package event

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateEvent(t *testing.T) {
	eventID := uuid.New()
	userID := uuid.New()
	venueID := uuid.New()
	events := []*Event{
		Reconstruct(
			eventID,
			"test_title",
			"test_description",
			time.Now().AddDate(0, 0, 1),
			time.Now().AddDate(0, 0, 2),
			time.Now(),
			time.Now(),
			userID,
			venueID,
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
			venueID,
		),
	}

	tests := []struct {
		name        string
		eventID     uuid.UUID
		title       string
		description string
		startDate   time.Time
		endDate     time.Time
		createdBy   uuid.UUID
		venueID     uuid.UUID
		mockFunc    func() *EventUpdater
		wantErr     bool
	}{
		{
			name:        "Update event successfully",
			eventID:     eventID,
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			createdBy:   userID,
			venueID:     venueID,
			mockFunc: func() *EventUpdater {
				var mockEventRepository MockEventRepository
				mockEventRepository.On("FetchByID", mock.Anything, mock.Anything).Return(events[0], nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return(events, nil)

				return NewEventUpdater(&mockEventRepository)
			},
			wantErr: false,
		},
		{
			name:        "Fetch event by id failed",
			eventID:     eventID,
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			createdBy:   userID,
			venueID:     venueID,
			mockFunc: func() *EventUpdater {
				var mockEventRepository MockEventRepository
				mockEventRepository.On("FetchByID", mock.Anything, mock.Anything).Return(nil, assert.AnError)

				return NewEventUpdater(&mockEventRepository)
			},
			wantErr: true,
		},
		{
			name:        "Event not found",
			eventID:     uuid.New(),
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			createdBy:   userID,
			venueID:     venueID,
			mockFunc: func() *EventUpdater {
				var mockEventRepository MockEventRepository
				mockEventRepository.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)

				return NewEventUpdater(&mockEventRepository)
			},
			wantErr: true,
		},
		{
			name:        "User not authorized",
			eventID:     eventID,
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			createdBy:   uuid.New(),
			venueID:     venueID,
			mockFunc: func() *EventUpdater {
				var mockEventRepository MockEventRepository
				mockEventRepository.On("FetchByID", mock.Anything, mock.Anything).Return(events[0], nil)

				return NewEventUpdater(&mockEventRepository)
			},
			wantErr: true,
		},
		{
			name:        "Fetch events by venue ID failed",
			eventID:     eventID,
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			createdBy:   userID,
			venueID:     venueID,
			mockFunc: func() *EventUpdater {
				var mockEventRepository MockEventRepository
				mockEventRepository.On("FetchByID", mock.Anything, mock.Anything).Return(events[0], nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return(nil, assert.AnError)

				return NewEventUpdater(&mockEventRepository)
			},
			wantErr: true,
		},
		{
			name:        "Set event start date after end date",
			eventID:     eventID,
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 5),
			endDate:     time.Now().AddDate(0, 0, 4),
			createdBy:   userID,
			venueID:     venueID,
			mockFunc: func() *EventUpdater {
				var mockEventRepository MockEventRepository
				mockEventRepository.On("FetchByID", mock.Anything, mock.Anything).Return(events[0], nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return(events, nil)

				return NewEventUpdater(&mockEventRepository)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			ctx := e.NewContext(nil, nil)

			updater := tt.mockFunc()
			event, err := updater.UpdateEvent(
				ctx,
				tt.eventID,
				tt.title,
				tt.description,
				tt.startDate,
				tt.endDate,
				tt.createdBy,
				tt.venueID,
			)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, event)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, event)
			assert.Equal(t, tt.title, event.Title())
			assert.Equal(t, tt.description, event.Description())
			assert.Equal(t, tt.startDate, event.StartDate())
			assert.Equal(t, tt.endDate, event.EndDate())
			assert.Equal(t, tt.createdBy, event.CreatedBy())
			assert.Equal(t, tt.venueID, event.VenueID())
		})
	}
}
