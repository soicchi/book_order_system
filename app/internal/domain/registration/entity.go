package registration

import (
	"time"

	"event_system/internal/domain/ticket"
	"event_system/internal/domain/status"

	"github.com/google/uuid"
)

type Registration struct {
	id uuid.UUID
	registeredAt time.Time
	status *status.Status
	ticket *ticket.Ticket 
}

func New(ticket *ticket.Ticket) *Registration {
	return &Registration{
		id: uuid.New(),
		registeredAt: time.Now(),
		status: status.New(status.Registered),
		ticket: ticket,
	}
}

func Reconstruct(
	id uuid.UUID,
	registeredAt time.Time,
	status *status.Status,
	ticket *ticket.Ticket,
) *Registration {
	return &Registration{
		id: id,
		registeredAt: registeredAt,
		status: status,
		ticket: ticket,
	}
}

func (r *Registration) ID() uuid.UUID {
	return r.id
}

func (r *Registration) RegisteredAt() time.Time {
	return r.registeredAt
}

func (r *Registration) Status() *status.Status {
	return r.status
}

func (r *Registration) Ticket() *ticket.Ticket {
	return r.ticket
}
