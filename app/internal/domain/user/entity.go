package user

import (
	"time"

	"event_system/internal/domain/role"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	name      string
	email     string
	role      role.Role
	createdAt time.Time
	updatedAt time.Time
}

func New(name, email string, role role.Role) *User {
	return &User{
		id:        uuid.New(),
		name:      name,
		email:     email,
		role:      role,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func Reconstruct(id uuid.UUID, name, email string, role role.Role, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		name:      name,
		email:     email,
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Role() role.Role {
	return u.role
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
