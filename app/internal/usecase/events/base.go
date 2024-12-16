package events

import (
	"event_system/internal/domain/event"
)

type EventUseCase struct {
	eventFactory    event.EventFactoryService
	eventUpdater    event.EventUpdaterService
	eventRepository event.EventRepository
}

func NewEventUseCase(
	eventFactory event.EventFactoryService,
	eventUpdater event.EventUpdaterService,
	eventRepository event.EventRepository,
) *EventUseCase {
	return &EventUseCase{
		eventFactory:    eventFactory,
		eventUpdater:    eventUpdater,
		eventRepository: eventRepository,
	}
}
