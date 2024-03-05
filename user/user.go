package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	UID     int    `json:"uid"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User = User{UID: 1, Name: "Brandon", Message: "Hello, World! 🐳"}
	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}