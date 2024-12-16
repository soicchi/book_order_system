package events

import (
	"testing"
	"time"

	"event_system/internal/domain/event"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateEvent(t *testing.T) {
	testEvent := event.Reconstruct(
		uuid.New(),
		"test_title",
		"test_description",
		time.Now().AddDate(0, 0, 1),
		time.Now().AddDate(0, 0, 2),
		time.Now(),
		time.Now(),
		uuid.New(),
		uuid.New(),
	)

	tests := []struct {
		name     string
		input    *UpdateInput
		mockFunc func() *EventUseCase
		wantErr  bool
	}{
		{
			name:  "Update event successfully",
			input: NewUpdateInput(uuid.New(), "title", "description", time.Now(), time.Now(), uuid.New(), uuid.New()),
			mockFunc: func() *EventUseCase {
				var mockEventUpdater event.MockEventUpdater
				var mockEventFactory event.MockEventFactory
				var mockEventRepository event.MockEventRepository
				mockEventUpdater.On(
					"UpdateEvent",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(testEvent, nil)
				mockEventRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

				return NewEventUseCase(&mockEventFactory, &mockEventUpdater, &mockEventRepository)
			},
			wantErr: false,
		},
		{
			name:  "Update event service failed",
			input: NewUpdateInput(uuid.New(), "title", "description", time.Now(), time.Now(), uuid.New(), uuid.New()),
			mockFunc: func() *EventUseCase {
				var mockEventUpdater event.MockEventUpdater
				var mockEventFactory event.MockEventFactory
				var mockEventRepository event.MockEventRepository
				mockEventUpdater.On(
					"UpdateEvent",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil, assert.AnError)

				return NewEventUseCase(&mockEventFactory, &mockEventUpdater, &mockEventRepository)
			},
			wantErr: true,
		},
		{
			name:  "Update event repository failed",
			input: NewUpdateInput(uuid.New(), "title", "description", time.Now(), time.Now(), uuid.New(), uuid.New()),
			mockFunc: func() *EventUseCase {
				var mockEventUpdater event.MockEventUpdater
				var mockEventFactory event.MockEventFactory
				var mockEventRepository event.MockEventRepository
				mockEventUpdater.On(
					"UpdateEvent",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(testEvent, nil)
				mockEventRepository.On("Update", mock.Anything, mock.Anything).Return(assert.AnError)

				return NewEventUseCase(&mockEventFactory, &mockEventUpdater, &mockEventRepository)
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

			usecase := tt.mockFunc()
			err := usecase.UpdateEvent(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
