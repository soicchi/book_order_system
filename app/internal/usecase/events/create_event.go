package events

import (
	"github.com/labstack/echo/v4"
)

func (eu *EventUseCase) CreateEvent(ctx echo.Context, input *CreateInput) error {
	event, err := eu.eventFactory.NewEvent(
		ctx,
		input.Title,
		input.Description,
		input.StartDate,
		input.EndDate,
		input.CreatedBy,
		input.VenueID,
	)
	if err != nil {
		return err
	}

	return eu.eventRepository.Create(ctx, event)
}
