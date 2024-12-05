package status

type StatusValue int

const (
	Registered StatusValue = iota + 1
	Canceled
)

func (s StatusValue) String() string {
	switch s {
	case Registered:
		return "Registered"
	case Canceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

func FromString(s string) StatusValue {
	switch s {
	case "Registered":
		return Registered
	case "Canceled":
		return Canceled
	default:
		return 0
	}
}

type Status struct {
	status StatusValue
}

func New(status StatusValue) Status {
	return Status{status: status}
}

func Reconstruct(status StatusValue) Status {
	return Status{status: status}
}

func (s Status) Value() StatusValue {
	return s.status
}
