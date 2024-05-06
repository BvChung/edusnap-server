package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ocr/cmd/response"
)

func JWTHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("cryptokey"), nil
		})

		if err != nil || !token.Valid {
			response.NewErrorResponse(w, "Invalid JWT token", response.InvalidRequest, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
