package event

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventRepository interface {
	Create(ctx echo.Context, event *Event) error
	FetchByID(ctx echo.Context, eventID uuid.UUID) (*Event, error)
	FetchAll(ctx echo.Context) ([]*Event, error)
	FetchByEventID(ctx echo.Context, eventID uuid.UUID) ([]*Event, error)
	Update(ctx echo.Context, event *Event) error
}
