package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	"github.com/google/uuid"
	"github.com/supabase-community/supabase-go"
)

type Student struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name string     `json:"name"`
}

type StudentHandler struct {
	DBClient       *supabase.Client
	VertexAIClient *genai.Client
}

func (sh *StudentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := sh.DBClient

	var students []Student

	_, err := client.From("student").Select("*", "", false).ExecuteTo(&students)

	if err != nil {
		fmt.Fprintf(w, "Cannot fetch data")
		return
	}

	if err := json.NewEncoder(w).Encode(students); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewStudentHandler(s *supabase.Client, v *genai.Client) *StudentHandler {
	return &StudentHandler{DBClient: s}
}
