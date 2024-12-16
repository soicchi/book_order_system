package events

import (
	"github.com/labstack/echo/v4"
)

func (eu *EventUseCase) UpdateEvent(ctx echo.Context, input *UpdateInput) error {
	event, err := eu.eventUpdater.UpdateEvent(
		ctx,
		input.EventID,
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

	return eu.eventRepository.Update(ctx, event)
}
