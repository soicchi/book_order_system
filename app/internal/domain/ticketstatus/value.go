package ticketstatus

type TicketStatusValue int

const (
	Active TicketStatusValue = iota + 1
	Used
	Canceled
)

func (t TicketStatusValue) String() string {
	switch t {
	case Active:
		return "Active"
	case Used:
		return "Used"
	case Canceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

func FromString(s string) TicketStatusValue {
	switch s {
	case "Active":
		return Active
	case "Used":
		return Used
	case "Canceled":
		return Canceled
	default:
		return 0
	}
}

type TicketStatus struct {
	status TicketStatusValue
}

func New(status TicketStatusValue) TicketStatus {
	return TicketStatus{status: status}
}

func Reconstruct(status TicketStatusValue) TicketStatus {
	return TicketStatus{status: status}
}

func (t TicketStatus) Value() TicketStatusValue {
	return t.status
}
