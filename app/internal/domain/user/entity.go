package user

import (
	"event_system/internal/domain/role"

	"github.com/google/uuid"
)

type User struct {
	id    uuid.UUID
	name  string
	email string
	role  *role.Role
}

func New(name, email string, role *role.Role) *User {
	return &User{
		id:    uuid.New(),
		name:  name,
		email: email,
		role:  role,
	}
}

func Reconstruct(id uuid.UUID, name, email string, role *role.Role) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
		role:  role,
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

func (u *User) Role() *role.Role {
	return u.role
}
