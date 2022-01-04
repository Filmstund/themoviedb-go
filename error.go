package themoviedb

import (
	"encoding/json"
	"fmt"
	"io"
)

type APIError struct {
	error   string
	errResp *ErrorResponse
}

func (apiErr *APIError) Error() string {
	return apiErr.error
}

func (apiErr *APIError) Response() *ErrorResponse {
	return apiErr.errResp
}

func newAPIError(body io.Reader, format string, a ...interface{}) *APIError {
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
