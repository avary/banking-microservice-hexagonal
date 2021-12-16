package errs

import "net/http"

type AppError struct {
	StatusCode int    `json:",omitempty"`
	Message    string `json:"message"`
}

func newAppError(message string, statusCode int) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewNotFoundError(message string) *AppError {
	return newAppError(message, http.StatusNotFound)
}

func NewUnexpectedError(message string) *AppError {
	return newAppError(message, http.StatusInternalServerError)
}
