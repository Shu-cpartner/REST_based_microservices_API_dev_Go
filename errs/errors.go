package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(messega string) *AppError {
	return &AppError{
		Message: messega,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(messega string) *AppError {
	return &AppError{
		Message: messega,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(messega string) *AppError {
	return &AppError{
		Message: messega,
		Code:    http.StatusUnprocessableEntity,
	}
}
