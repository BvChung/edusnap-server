package middleware

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

type responseLoggerWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseLoggerWrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		statusLogger := &responseLoggerWrapper{w, http.StatusOK}

		log.Printf("[ID: %s] Incoming %s request to %s", id, r.Method, r.URL.Path)

		next.ServeHTTP(statusLogger, r)

		log.Printf("[ID: %s] Completed %s request to %s with status %d", id, r.Method, r.URL.Path, statusLogger.statusCode)
	})
}
