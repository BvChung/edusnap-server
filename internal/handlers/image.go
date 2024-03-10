package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/internal/supabase"
)

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
	res, err := client.From("student").Select("*", "", false).ExecuteTo(&students)

	if err != nil {
		fmt.Fprintf(w, "Cannot fetch data")
		return
	}

	fmt.Println(res)

	if err := json.NewEncoder(w).Encode(students); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}