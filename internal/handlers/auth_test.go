package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func Test_LoginHandler(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal("Unable to load env")
	}

	data := []struct{
		Email string
		Password string
		ExpectedStatusCode int
	}{
		{Email: "", Password: "", ExpectedStatusCode: 500},
		{Email: "", Password: "pw", ExpectedStatusCode: 500},
		{Email: "b@gmail.com", Password: "", ExpectedStatusCode: 500},
		{Email: "b@gmail.com", Password: "pw", ExpectedStatusCode: 200},
		{Email: "asff@gmail.com", Password: "pw", ExpectedStatusCode: 401},
	}

	for _, d := range data {
		t.Run("Test login route", func(t *testing.T) {
			user := &LoginUser{
				Email:    d.Email,
				Password: d.Password,
			}
		
			jsonData, err := json.Marshal(user)
			if err != nil {
				t.Fatal(err)
			}

			r, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
	
			if err != nil {
				t.Fatal(err.Error())
			}
	
			r.Header.Set("Content-Type", "application/json")
		
			w := httptest.NewRecorder()
	
			handler := http.HandlerFunc(LoginHandler)
	
			handler.ServeHTTP(w, r)
	
			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}