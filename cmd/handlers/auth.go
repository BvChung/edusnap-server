package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/cmd/response"
	"github.com/ocr/cmd/services/encrypt"
	"github.com/supabase-community/supabase-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       *uuid.UUID `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}

type ReturnedUser struct {
	ID    *uuid.UUID `json:"id"`
	Email string     `json:"email"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginHandler struct {
	DBClient *supabase.Client
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "POST":
		LoginUser(lh.DBClient, w, r)
	}
}

type RegisterHandler struct {
	DBClient *supabase.Client
}

func (rh *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "POST":
		RegisterUser(rh.DBClient, w, r)
	}
}

func NewLoginHandler(s *supabase.Client) *LoginHandler {
	return &LoginHandler{DBClient: s}
}

func NewRegisterHandler(s *supabase.Client) *RegisterHandler {

	return &RegisterHandler{DBClient: s}
}

// func () NewLoginHandler() http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 	})
// }

func LoginUser(client *supabase.Client, w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundUser User

	if _, err := client.From("users").Select("*", "exact", false).Eq("email", credentials.Email).Single().ExecuteTo(&foundUser); err != nil {
		log.Println("Email not found:", err.Error())
		response.NewErrorResponse(w, "Email not found", response.InvalidRequest, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password)); err != nil {
		log.Println("Invalid password:", err.Error())
		response.NewErrorResponse(w, "Invalid password", response.InvalidRequest, http.StatusUnauthorized)
		return
	}

	var returnedUser ReturnedUser = ReturnedUser{ID: foundUser.ID, Email: foundUser.Email}
	data := []ReturnedUser{returnedUser}

	response.NewSuccessResponse(w, data, http.StatusOK)
}

func RegisterUser(client *supabase.Client, w http.ResponseWriter, r *http.Request) {
	var credentials User

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		response.NewErrorResponse(w, "Internal server error", response.ServerError, http.StatusInternalServerError)
		return
	}

	hashedPw, err := encrypt.HashPassword(credentials.Password)

	if err != nil {
		log.Println(err.Error())
		response.NewErrorResponse(w, "Password cannot be empty", response.InvalidRequest, http.StatusBadRequest)
		return
	}

	var data []ReturnedUser
	var row Credentials = Credentials{Email: credentials.Email, Password: string(hashedPw)}

	if _, err := client.From("users").Insert(row, false, "", "", "").ExecuteTo(&data); err != nil {
		log.Println("User with email already exists:", err.Error())
		response.NewErrorResponse(w, "User with email already exists", response.InvalidRequest, http.StatusConflict)
		return
	}

	response.NewSuccessResponse(w, data, http.StatusCreated)
}
