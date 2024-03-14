package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	InvalidRequest = "invalid_request_error"
	ServerError = "server_error"
)

type SuccessResponse[T any] struct {
	Data []T `json:"data"`
}

type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func NewSuccessResponse[T any](w http.ResponseWriter, data []T, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	res := &SuccessResponse[T]{Data: data}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return fmt.Errorf("failed to encode data to JSON: %w", err)
	}

	return nil
}

func NewErrorResponse(w http.ResponseWriter, message string, errorCode string, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(statusCode)

	res := ErrorResponse{
		Error: ErrorDetails{Code: errorCode, Message: message},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return fmt.Errorf("failed to encode data to JSON: %w", err)
	}

	return nil
}