package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("WWTm00/88EzBgk5y7IlcpZDIDDUP2vtA/EuCDmujo5j2pD0ZjRll/TaLlYoE4KlQUg7LsSMUFAU278nSrF/Jhg==")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Unauthorized no authorization header")
			return
		}

		jwtToken := strings.Split(authHeader, " ")[1]

		_, err := validateToken(jwtToken)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "invalid jwt token")
			return
		}

		// fmt.Println(t.Claims)
		// fmt.Println(t.Claims.GetIssuer())

		next.ServeHTTP(w, r)
	})
}

func createToken() (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "ocr-app",                        // Issuer
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenStr, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func validateToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return t, nil
}
