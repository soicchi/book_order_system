package venue

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type VenueRepository interface {
	Create(ctx echo.Context, v *Venue) error
	FetchAll(ctx echo.Context) ([]*Venue, error)
	FetchByID(ctx echo.Context, venueID uuid.UUID) (*Venue, error)
	Update(ctx echo.Context, v *Venue) error
}
