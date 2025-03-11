package errors

import (
	"errors"
	"net/http"
)

type Errors struct {
	StatusCode int
	Code       string            `json:"code"`
	Message    string            `json:"message"`
	Details    map[string]string `json:"details"`
}

func New(code int, message string, details map[string]string) *Errors {
	return &Errors{
		StatusCode: code,
		Code:       http.StatusText(code),
		Message:    message,
		Details:    details,
	}
}

func (e Errors) Error() string {
	return e.Message
}

func ToErrors(err error) *Errors {
	if errors.Is(err, &Errors{}) {
		return err.(*Errors)
	}
	return nil
}
