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

func TestCreateEvent(t *testing.T) {
	tests := []struct {
		name     string
		dto      *CreateInput
		mockFunc func() *EventUseCase
		wantErr  bool
	}{
		{
			name: "Create event successfully",
			dto:  NewCreateInput("title", "description", time.Now(), time.Now(), uuid.New(), uuid.New()),
			mockFunc: func() *EventUseCase {
				var mockEventFactory event.MockEventFactory
				var mockEventUpdater event.MockEventUpdater
				var mockEventRepository event.MockEventRepository
				mockEventFactory.On(
					"NewEvent",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(&event.Event{}, nil)
				mockEventRepository.On("Create", mock.Anything, mock.Anything).Return(nil)

				return NewEventUseCase(&mockEventFactory, &mockEventUpdater, &mockEventRepository)
			},
			wantErr: false,
		},
		{
			name: "Create event failed",
			dto:  NewCreateInput("title", "description", time.Now(), time.Now(), uuid.New(), uuid.New()),
			mockFunc: func() *EventUseCase {
				var mockEventFactory event.MockEventFactory
				var mockEventUpdater event.MockEventUpdater
				var mockEventRepository event.MockEventRepository
				mockEventFactory.On(
					"NewEvent",
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
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			ctx := e.NewContext(nil, nil)

			uc := tt.mockFunc()
			err := uc.CreateEvent(ctx, tt.dto)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
