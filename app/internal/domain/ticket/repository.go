package ticket

import (
	"github.com/labstack/echo/v4"
)

type TicketRepository interface {
	Create(ctx echo.Context, ticket *Ticket) error
	FetchByQRCode(ctx echo.Context, qrCode string) (*Ticket, error)
	Update(ctx echo.Context, ticket *Ticket) error
}
