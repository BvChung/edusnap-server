package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ocr/cmd/database"
)

func Test_LoginHandler(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal("Unable to load env")
	}

	data := []struct {
		Email              string
		Password           string
		ExpectedStatusCode int
	}{
		{Email: "", Password: "", ExpectedStatusCode: 401},
		{Email: "", Password: "pw", ExpectedStatusCode: 401},
		{Email: "b@gmail.com", Password: "", ExpectedStatusCode: 401},
		{Email: "b@gmail.com", Password: "pw", ExpectedStatusCode: 200},
		{Email: "asff@gmail.com", Password: "pw", ExpectedStatusCode: 401},
	}

	client, err := database.CreateSupabaseClient()

	if err != nil {
		t.Fatalf("Unable to connect to database: %s", err.Error())
	}

	for _, d := range data {
		t.Run("Test login route", func(t *testing.T) {
			user := &Credentials{
				Email:    d.Email,
				Password: d.Password,
			}

			jsonData, err := json.Marshal(user)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()
			handler := http.Handler(NewLoginHandler(client))

			r.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(w, r)

			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, d.ExpectedStatusCode)
			}
		})
	}
}

func Test_RegisterHandler(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal("Unable to load env")
	}

	data := []struct {
		Email              string
		Password           string
		ExpectedStatusCode int
	}{
		{Email: "person@gmail.com", Password: "", ExpectedStatusCode: 400},
		{Email: "b@gmail.com", Password: "pw", ExpectedStatusCode: 409},
	}

	client, err := database.CreateSupabaseClient()

	if err != nil {
		t.Fatalf("Unable to connect to database: %s", err.Error())
	}

	for _, d := range data {
		t.Run("Test register route", func(t *testing.T) {
			user := &Credentials{
				Email:    d.Email,
				Password: d.Password,
			}

			jsonData, err := json.Marshal(user)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()
			handler := http.Handler(NewRegisterHandler(client))

			r.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(w, r)

			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, d.ExpectedStatusCode)
			}
		})
	}
}
