package format

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SuccessResponse struct {
    Data interface{} `json:"data"`
}

type ErrorResponse struct {
    Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Type string `json:"type"`
}

func NewSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) error {
    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(statusCode)

    res := &SuccessResponse{Data: data}

    if err := json.NewEncoder(w).Encode(res); err != nil {
        return fmt.Errorf("failed to encode data to JSON: %w", err)
    }

	return nil
}

func NewErrorResponse(w http.ResponseWriter, message string, errorCode string, statusCode int) error {
    w.Header().Set("Content-Type", "application/json")
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(statusCode)

    res := ErrorResponse{
        Error: ErrorDetails{Code: errorCode, Message: message},
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