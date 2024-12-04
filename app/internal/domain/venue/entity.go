package venue

import (
	"github.com/google/uuid"
)

type Venue struct {
	id          uuid.UUID
	name        string
	address string
	capacity int
}

func New(name, address string, capacity int) *Venue {
	return &Venue{
		id:      uuid.New(),
		name:    name,
		address: address,
		capacity: capacity,
	}
}

func Reconstruct(id uuid.UUID, name, address string, capacity int) *Venue {
	return &Venue{
		id:      id,
		name:    name,
		address: address,
		capacity: capacity,
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
