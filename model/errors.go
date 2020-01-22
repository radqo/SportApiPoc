package model

// APIError Error returned by api call or internal server error 500
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}