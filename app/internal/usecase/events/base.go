package events

import (
	"event_system/internal/domain/event"
)

type EventUseCase struct {
	eventFactory    event.EventFactoryService
	eventRepository event.EventRepository
}

func NewEventUseCase(eventFactory event.EventFactoryService, eventRepository event.EventRepository) *EventUseCase {
	return &EventUseCase{
		eventFactory:    eventFactory,
		eventRepository: eventRepository,
	}
}
