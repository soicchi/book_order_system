package errors

type ErrorIssue int

// Add ErrorIssue when you need a new error issue
const (
	NoIssue ErrorIssue = iota
	Required
	Invalid
	ZeroOrLess
	LessThanZero
	Empty
	Unknown
)

func (e ErrorIssue) String() string {
	switch e {
	case Required:
		return "Required"
	case Invalid:
		return "Invalid"
	case ZeroOrLess:
		return "ZeroOrLess"
	case LessThanZero:
		return "LessThanZero"
	case Empty:
		return "Empty"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}
