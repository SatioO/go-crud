package helpers

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	StatusCode int `json:"status_code"`
	Msg        any `json:"message"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("error: %d", e.StatusCode)
}

func NewApiError(statusCode int, err error) ApiError {
	return ApiError{StatusCode: statusCode, Msg: err.Error()}
}

func InvalidRequestData(errors map[string]string) ApiError {
	return ApiError{StatusCode: http.StatusUnprocessableEntity, Msg: errors}
}

func InvalidJson() ApiError {
	return ApiError{StatusCode: http.StatusBadRequest, Msg: fmt.Errorf("invalid JSON request data")}
}
