package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	response "github.com/ocr/internal/format"
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

func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("unable to hash password, %w", err) 
	}

	return hash, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials User
	
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	client, err := supabase.CreateClient()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundUser User;
	
	if _ ,err := client.From("users").Select("*", "exact", false).Eq("email", credentials.Email).Single().ExecuteTo(&foundUser); err != nil {
		log.Println("Email not found: ", err.Error())
		response.NewErrorResponse(w, "Email not found", "INVALID_CREDENTIALS", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password)); err != nil {
		log.Println("Invalid password: ", err.Error())
		response.NewErrorResponse(w, "Invalid password", "INVALID_CREDENTIALS", http.StatusUnauthorized)
		return
	}

	var returnedUser ReturnedUser = ReturnedUser{ID: foundUser.ID, Email: foundUser.Email}
	
	response.NewSuccessResponse(w, &returnedUser, http.StatusOK)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		response.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	client, err := supabase.CreateClient()

	if err != nil {
		log.Println(err.Error())
		response.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	hashedPw, err := hashPassword(newUser.Password)

	if err != nil {
		log.Println(err.Error())
		response.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	var row RegisterUser = RegisterUser{Email: newUser.Email, Password: string(hashedPw)}

	if _, _, err := client.From("users").Insert(row, false, "", "", "").Execute(); err != nil {
		log.Println("User with email already exists", err.Error())
		response.NewErrorResponse(w, "User with email already exists", "INVALID_CREDENTIALS", http.StatusConflict)
		return
	}

	response.NewSuccessResponse(w, map[string]string{
		"message": "User registered successfully",
		"email":   newUser.Email,
	}, http.StatusCreated)
}