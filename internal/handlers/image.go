package handlers

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	"github.com/google/uuid"
	"github.com/ocr/internal/response"
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

	response.NewSuccessResponse(w, students, http.StatusOK)
}

func NewStudentHandler(s *supabase.Client, v *genai.Client) *StudentHandler {
	return &StudentHandler{DBClient: s, VertexAIClient: v}
}

// mux.HandleFunc("GET /profile/{id}", func(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")

// 	search := r.URL.Query().Get("search")
// 	fmt.Println(search)
// 	fmt.Fprintf(w, "Parameter: %s", id)
// })
