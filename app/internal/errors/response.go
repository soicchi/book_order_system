package errors

type ErrorDetails struct {
	Identifiers map[string]interface{} `json:"identifier,omitempty"`
	Resource    string                 `json:"resource"`
}

type ErrorResponse struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Details ErrorDetails `json:"details"`
}

type Identifiers map[string]interface{}

type option func(*ErrorResponse)

func WithIdentifiers(identifiers Identifiers) option {
	return func(e *ErrorResponse) {
		e.Details.Identifiers = identifiers
	}
}

func newErrorResponse(code string, message, resource string, options ...option) ErrorResponse {
	res := ErrorResponse{
		Code:    code,
		Message: message,
	}

	for _, opt := range options {
		opt(&res)
	}

	return res
}
