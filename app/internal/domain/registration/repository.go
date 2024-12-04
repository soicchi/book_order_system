package registration

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RegistrationRepository interface {
	Create(ctx echo.Context, registration *Registration) error
	Update(ctx echo.Context, registration *Registration) error
	FetchByEventID(ctx echo.Context, eventID uuid.UUID) ([]*Registration, error)
}
