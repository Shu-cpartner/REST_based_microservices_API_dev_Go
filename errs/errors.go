package errs

import (
	"net/http"
)

type AppError struct {
	Message string
	Code    int
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
