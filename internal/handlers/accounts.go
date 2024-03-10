package handlers

import (
	"log"
	"net/http"

	"github.com/ocr/internal/format"
	"github.com/ocr/internal/supabase"
)

func AccountsHandler(w http.ResponseWriter, r *http.Request) {
    client, err := supabase.CreateClient()

	if err != nil {
		log.Println(err.Error())
		format.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	var data []ReturnedUser
	if _, err := client.From("users").Select("*", "", false).Eq("id", "").ExecuteTo(&data); err != nil {
		format.NewErrorResponse(w, "Cannot find user", "NOT_FOUND", http.StatusNotFound)
		return 
	}

	format.NewSuccessResponse(w, data, http.StatusOK)
}