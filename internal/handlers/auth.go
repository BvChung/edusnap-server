package handlers

import (
	"encoding/json"	
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/internal/encrypt"
	"github.com/ocr/internal/format"
	"github.com/ocr/internal/database"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID      *uuid.UUID  `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ReturnedUser struct {
	ID      *uuid.UUID  `json:"id"`
	Email string `json:"email"`
}

type RegisterUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials LoginUser
	
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	client, err := database.CreateSupabaseClient()

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundUser User;
	
	if _ ,err := client.From("users").Select("*", "exact", false).Eq("email", credentials.Email).Single().ExecuteTo(&foundUser); err != nil {
		log.Println("Email not found:", err.Error())
		format.NewErrorResponse(w, "Email not found", "INVALID_CREDENTIALS", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password)); err != nil {
		log.Println("Invalid password:", err.Error())
		format.NewErrorResponse(w, "Invalid password", "INVALID_CREDENTIALS", http.StatusUnauthorized)
		return
	}

	var returnedUser ReturnedUser = ReturnedUser{ID: foundUser.ID, Email: foundUser.Email} 
	data := []ReturnedUser{returnedUser}
	
	format.NewSuccessResponse(w, data, http.StatusOK)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var credentials User

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		format.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	client, err := database.CreateSupabaseClient()

	if err != nil {
		log.Println(err.Error())
		format.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	hashedPw, err := encrypt.HashPassword(credentials.Password)

	if err != nil {
		log.Println(err.Error())
		format.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	var data []ReturnedUser
	var row RegisterUser = RegisterUser{Email: credentials.Email, Password: string(hashedPw)}

	if _, err := client.From("users").Insert(row, false, "", "", "").ExecuteTo(&data); err != nil {
		log.Println("User with email already exists:", err.Error())
		format.NewErrorResponse(w, "User with email already exists", "EMAIL_ALREADY_IN_USE", http.StatusConflict)
		return
	}

	format.NewSuccessResponse(w, data, http.StatusCreated)
}