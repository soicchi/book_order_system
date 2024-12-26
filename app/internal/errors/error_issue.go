package errors

type ErrorIssue int

// Add ErrorIssue when you need a new error issue
const (
	Unknown ErrorIssue = iota
	Required
	Invalid
	ZeroOrLess
	LessThanZero
	Empty
	InvalidTimeRange
	NotOrganizer
	NotCreator
	ExceedCapacity
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
	case InvalidTimeRange:
		return "InvalidTimeRange"
	case NotOrganizer:
		return "NotOrganizer"
	case NotCreator:
		return "NotCreator"
	case ExceedCapacity:
		return "ExceedCapacity"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}
