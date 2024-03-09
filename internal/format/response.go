package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
    Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
    Code    int16 `json:"code"`
    Message string `json:"message"`
    Status string `json:"status"`
}

func NewSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) error {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(statusCode)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        return fmt.Errorf("failed to encode data to JSON: %w", err)
    }

	return nil
}

func NewErrorResponse(w http.ResponseWriter, message string, status string, statusCode int) error {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(statusCode)

    res := ErrorResponse{
        Error: ErrorDetails{Code: int16(statusCode), Message: message, Status: status},
    }

    if err := json.NewEncoder(w).Encode(res); err != nil {
        return fmt.Errorf("failed to encode data to JSON: %w", err)
    }

	return nil
}

func CountWrapper(f http.HandlerFunc) http.HandlerFunc {
	timesCalled := 0

    return func(w http.ResponseWriter, r *http.Request) {
        // Pre-processing: Add your code here
		timesCalled++
		log.Printf("Times called: %d \n", timesCalled)

        // Call the original handler
        f(w, r)

        // Post-processing: Add your code here
    }
}