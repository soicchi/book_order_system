package registration

import (
	"time"

	"event_system/internal/domain/status"

	"github.com/google/uuid"
)

type Registration struct {
	id           uuid.UUID
	registeredAt time.Time
	status       status.Status
	eventID      uuid.UUID
	userID       uuid.UUID
}

func New(eventID, userID uuid.UUID) *Registration {
	return &Registration{
		id:           uuid.New(),
		registeredAt: time.Now(),
		status:       status.New(status.Registered),
		eventID:      eventID,
		userID:       userID,
	}
}

func Reconstruct(
	id uuid.UUID,
	registeredAt time.Time,
	status status.Status,
	eventID uuid.UUID,
	userID uuid.UUID,
) *Registration {
	return &Registration{
		id:           id,
		registeredAt: registeredAt,
		status:       status,
		eventID:      eventID,
		userID:       userID,
	}
}

func (r *Registration) ID() uuid.UUID {
	return r.id
}

func (r *Registration) RegisteredAt() time.Time {
	return r.registeredAt
}

func (r *Registration) Status() status.Status {
	return r.status
}

func (r *Registration) EventID() uuid.UUID {
	return r.eventID
}

func (r *Registration) UserID() uuid.UUID {
	return r.userID
}
