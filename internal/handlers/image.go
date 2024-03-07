package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/internal/supabase"
)

type Profile struct {
	ID      *uuid.UUID  `json:"id,omitempty"`
	Username string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type Student struct {
	ID      *uuid.UUID  `json:"id,omitempty"`
	Name string `json:"name"`
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	client, clientErr := supabase.CreateClient()

	if clientErr != nil {
		fmt.Fprintf(w, "Could not connect to db")
		return
	}

	var students []Student
	err := client.DB.From("student").Select("*").Execute(&students)

	if err != nil {
		fmt.Fprintf(w, "Cannot fetch data")
		return
	}

	if err := json.NewEncoder(w).Encode(students); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	client, clientErr := supabase.CreateClient()

	if clientErr != nil {
		fmt.Fprintf(w, "Could not connect to db")
		return
	}

	var profile []Profile
	err := client.DB.From("profiles").Select("id, username, avatar").Execute(&profile)

	if err != nil {
		fmt.Fprintf(w, "Cannot fetch data")
		return
	}


	if err := json.NewEncoder(w).Encode(profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}