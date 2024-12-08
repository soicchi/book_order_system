package role

type RoleValue int

const (
	Organizer RoleValue = iota + 1
	Attendee
)

func (r RoleValue) String() string {
	switch r {
	case Organizer:
		return "organizer"
	case Attendee:
		return "attendee"
	default:
		return "unknown"
	}
}

func FromString(s string) RoleValue {
	switch s {
	case "Organizer":
		return Organizer
	case "Attendee":
		return Attendee
	default:
		return 0
	}
}

type Role struct {
	role RoleValue
}

func New(role RoleValue) Role {
	return Role{role: role}
}

func Reconstruct(role RoleValue) Role {
	return Role{role: role}
}

func (r Role) Value() RoleValue {
	return r.role
}
