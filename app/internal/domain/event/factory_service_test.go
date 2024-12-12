package event

import (
	"testing"
	"time"

	"event_system/internal/domain/role"
	"event_system/internal/domain/user"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventFactoryNew(t *testing.T) {
	testOrganizer := user.New("organizer", "organizer@test.com", role.New(role.Organizer))
	testAttendee := user.New("attendee", "attendee@test.com", role.New(role.Attendee))
	testEvents := []*Event{
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
	}

	tests := []struct {
		name        string
		mockFunc    func() *EventFactory
		title       string
		description string
		startDate   time.Time
		endDate     time.Time
		userID      uuid.UUID
		venueID     uuid.UUID
		wantErr     bool
	}{
		{
			name: "Create event entity successfully",
			mockFunc: func() *EventFactory {
				var mockUserRepository user.MockUserRepository
				var mockEventRepository MockEventRepository
				mockUserRepository.On("FetchByID", mock.Anything, mock.Anything).Return(testOrganizer, nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return(testEvents, nil)

				return NewEventFactory(&mockEventRepository, &mockUserRepository)
			},
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			userID:      testOrganizer.ID(),
			venueID:     uuid.New(),
			wantErr:     false,
		},
		{
			name: "User not found",
			mockFunc: func() *EventFactory {
				var mockUserRepository user.MockUserRepository
				var mockEventRepository MockEventRepository
				mockUserRepository.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)

				return NewEventFactory(&mockEventRepository, &mockUserRepository)
			},
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			userID:      testOrganizer.ID(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name: "User is not an organizer",
			mockFunc: func() *EventFactory {
				var mockUserRepository user.MockUserRepository
				var mockEventRepository MockEventRepository
				mockUserRepository.On("FetchByID", mock.Anything, mock.Anything).Return(testAttendee, nil)

				return NewEventFactory(&mockEventRepository, &mockUserRepository)
			},
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			userID:      testAttendee.ID(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name: "Fetch events by venue ID failed",
			mockFunc: func() *EventFactory {
				var mockUserRepository user.MockUserRepository
				var mockEventRepository MockEventRepository
				mockUserRepository.On("FetchByID", mock.Anything, mock.Anything).Return(testOrganizer, nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return([]*Event{}, assert.AnError)

				return NewEventFactory(&mockEventRepository, &mockUserRepository)
			},
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 3),
			endDate:     time.Now().AddDate(0, 0, 4),
			userID: 	testOrganizer.ID(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name: "Event already exists in the venue",
			mockFunc: func() *EventFactory {
				var mockUserRepository user.MockUserRepository
				var mockEventRepository MockEventRepository
				mockUserRepository.On("FetchByID", mock.Anything, mock.Anything).Return(testOrganizer, nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return(testEvents, nil)

				return NewEventFactory(&mockEventRepository, &mockUserRepository)
			},
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 1),
			endDate:     time.Now().AddDate(0, 0, 2),
			userID:      testOrganizer.ID(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
		{
			name: "Event end date is before start date",
			mockFunc: func() *EventFactory {
				var mockUserRepository user.MockUserRepository
				var mockEventRepository MockEventRepository
				mockUserRepository.On("FetchByID", mock.Anything, mock.Anything).Return(testOrganizer, nil)
				mockEventRepository.On("FetchByVenueID", mock.Anything, mock.Anything).Return([]*Event{}, nil)

				return NewEventFactory(&mockEventRepository, &mockUserRepository)
			},
			title:       "test_title",
			description: "test_description",
			startDate:   time.Now().AddDate(0, 0, 5),
			endDate:     time.Now().AddDate(0, 0, 4),
			userID:      testOrganizer.ID(),
			venueID:     uuid.New(),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			ctx := e.NewContext(nil, nil)

			eventFactory := tt.mockFunc()

			event, err := eventFactory.New(
				ctx,
				tt.title,
				tt.description,
				tt.startDate,
				tt.endDate,
				tt.userID,
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
			assert.Equal(t, tt.userID, event.CreatedBy())
			assert.Equal(t, tt.venueID, event.VenueID())
		})
	}
}
