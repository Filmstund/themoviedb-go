package themoviedb

import (
	"encoding/json"
	"fmt"
	"io"
)

// APIError is returned when the API responds with a non-OK status.
type APIError struct {
	error   string
	errResp *ErrorResponse
}

// Error implements the error interface.
func (apiErr *APIError) Error() string {
	return apiErr.error
}

// Response returns the decoded error response from the API.
func (apiErr *APIError) Response() *ErrorResponse {
	return apiErr.errResp
}

func newAPIError(body io.Reader, format string, a ...any) *APIError {
	decoder := json.NewDecoder(body)
	errResp := new(ErrorResponse)
	msg := fmt.Sprintf(format, a...)
	if err := decoder.Decode(errResp); err != nil {
		msg += fmt.Sprintf(" (%v)", err)
	}
	return &APIError{
		error:   msg,
		errResp: errResp,
	}
}

// ErrorResponse describes errors returned from the API.
type ErrorResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}
