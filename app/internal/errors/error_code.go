package errors

import (
	"net/http"
)

type ErrorCode int

const (
	ValidationError ErrorCode = iota + 1
	AuthenticationError
	AuthorizationError
	NotFoundError
	AlreadyExistError
	UnexpectedError
)

func (e ErrorCode) String() string {
	switch e {
	case ValidationError:
		return "Validation"
	case AuthenticationError:
		return "Authentication"
	case AuthorizationError:
		return "Authorization"
	case NotFoundError:
		return "NotFound"
	case AlreadyExistError:
		return "AlreadyExist"
	case UnexpectedError:
		return "Unexpected"
	default:
		return "Unexpected"
	}
}

func (e ErrorCode) Status() int {
	switch e {
	case ValidationError:
		return http.StatusBadRequest
	case AuthenticationError:
		return http.StatusUnauthorized
	case AuthorizationError:
		return http.StatusForbidden
	case NotFoundError:
		return http.StatusNotFound
	case AlreadyExistError:
		return http.StatusConflict
	case UnexpectedError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (e ErrorCode) Message() string {
	switch e {
	case ValidationError:
		return "Validation error"
	case AuthenticationError:
		return "Authentication error"
	case AuthorizationError:
		return "Authorization error"
	case NotFoundError:
		return "Not found error"
	case AlreadyExistError:
		return "Already exist error"
	case UnexpectedError:
		return "Unexpected error"
	default:
		return "Unexpected error"
	}
}
