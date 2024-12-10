package ticket

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TicketRepository interface {
	Create(ctx echo.Context, ticket *Ticket) error
	FetchByRegistrationID(ctx echo.Context, registrationID uuid.UUID) (*Ticket, error)
	Update(ctx echo.Context, ticket *Ticket) error
}
