package errs

import (
	"encoding/json"
	"io"
	"net/http"
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode,omitempty"`
}

func (e *AppError) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (e *AppError) AsMessage() string {
	return e.Message
}

func NewNotFoundError(m string) *AppError {
	return &AppError{
		Message:    m,
		StatusCode: http.StatusNotFound,
	}
}

func NewUnexpectedError(m string) *AppError {
	return &AppError{
		Message:    m,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewInternalServerError(m string) *AppError {
	return &AppError{
		Message:    m,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}
