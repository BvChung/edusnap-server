package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/internal/supabase"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID      *uuid.UUID  `json:"id,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ReturnedUser struct {
	ID      *uuid.UUID  `json:"id,omitempty"`
	Email string `json:"email"`
}

type RegisterUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("unable to create password, %s", err.Error()) 
	}

	return hash, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials User
	
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pw, err := HashPassword(credentials.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	credentials.Password = string(pw)
	
	client, err := supabase.CreateClient()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundUser User;

	if err := client.DB.From("users").Select("*").Eq("email", credentials.Email).Execute(&foundUser); err != nil {
		log.Println(err.Error())
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	fmt.Println(foundUser)

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	var returnedUser ReturnedUser = ReturnedUser{ID: foundUser.ID, Email: foundUser.Email}
	
	if err := json.NewEncoder(w).Encode(returnedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, err := supabase.CreateClient()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPw, err := HashPassword(newUser.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []User
	var row RegisterUser = RegisterUser{Email: newUser.Email, Password: string(hashedPw)}

	if err := client.DB.From("users").Insert(row).Execute(&res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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


