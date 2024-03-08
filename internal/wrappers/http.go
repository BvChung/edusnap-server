package wrappers

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) error {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(statusCode)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        return fmt.Errorf("")
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