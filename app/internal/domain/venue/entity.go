package venue

import (
	"time"

	"github.com/google/uuid"
)

type Venue struct {
	id        uuid.UUID
	name      string
	address   string
	capacity  int
	createdAt time.Time
	updatedAt time.Time
}

func New(name, address string, capacity int) *Venue {
	return &Venue{
		id:        uuid.New(),
		name:      name,
		address:   address,
		capacity:  capacity,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func Reconstruct(id uuid.UUID, name, address string, capacity int, createdAt, updatedAt time.Time) *Venue {
	return &Venue{
		id:        id,
		name:      name,
		address:   address,
		capacity:  capacity,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (v *Venue) ID() uuid.UUID {
	return v.id
}

func (v *Venue) Name() string {
	return v.name
}

func (v *Venue) Address() string {
	return v.address
}

func (v *Venue) Capacity() int {
	return v.capacity
}

func (v *Venue) CreatedAt() time.Time {
	return v.createdAt
}

func (v *Venue) UpdatedAt() time.Time {
	return v.updatedAt
}
