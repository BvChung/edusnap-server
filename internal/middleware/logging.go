package middleware

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()

		log.Printf("[ID: %s] Incoming %s request from %s to route: %s", id, r.Method, r.RemoteAddr, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("[ID: %s] Completed %s request from %s to route: %s", id, r.Method, r.RemoteAddr, r.URL.Path)

	})
}
