package ticket

import (
	"time"

	"event_system/internal/domain/ticketstatus"

	"github.com/google/uuid"
)

type Ticket struct {
	id           uuid.UUID
	qrCode       string
	issuedAt     time.Time
	usedAt       time.Time
	ticketStatus *ticketstatus.TicketStatus
}

func New(qrCode string) *Ticket {
	return &Ticket{
		id:           uuid.New(),
		qrCode:       qrCode,
		issuedAt:     time.Now(),
		usedAt:       time.Time{},
		ticketStatus: ticketstatus.New(ticketstatus.Active),
	}
}

func Reconstruct(
	id uuid.UUID,
	qrCode string,
	issuedAt time.Time,
	usedAt time.Time,
	ticketStatus *ticketstatus.TicketStatus,
) *Ticket {
	return &Ticket{
		id:           id,
		qrCode:       qrCode,
		issuedAt:     issuedAt,
		usedAt:       usedAt,
		ticketStatus: ticketStatus,
	}
}

func (t *Ticket) ID() uuid.UUID {
	return t.id
}

func (t *Ticket) QRCode() string {
	return t.qrCode
}

func (t *Ticket) IssuedAt() time.Time {
	return t.issuedAt
}

func (t *Ticket) UsedAt() time.Time {
	return t.usedAt
}

func (t *Ticket) TicketStatus() *ticketstatus.TicketStatus {
	return t.ticketStatus
}
