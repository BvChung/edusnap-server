package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func Test_AccountsRoute(t *testing.T) {
	t.Run("Test accounts route", func(t *testing.T) {
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatal("Unable to load env")
		}

		r := httptest.NewRequest("GET", "/api/accounts", nil)
		w := httptest.NewRecorder()

		handler := http.Handler(http.HandlerFunc(AccountsHandler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}