package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	UID     int    `json:"uid"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	f := r.Form

	fmt.Println(f)

	var user User
	
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Message = "🐳🐳🐳🐳🐳"
	
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User = User{UID: 1, Name: "Brandon", Message: "Hello, World! 🐳"}
	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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


