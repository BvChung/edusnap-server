package handlers

import (
	"log"
	"net/http"

	"github.com/ocr/cmd/database"
	"github.com/ocr/cmd/response"
)

func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	client, err := database.CreateSupabaseClient()

	if err != nil {
		log.Println(err.Error())
		response.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	var data []ReturnedUser
	if _, err := client.From("users").Select("*", "", false).ExecuteTo(&data); err != nil {
		response.NewErrorResponse(w, "Cannot find user", "NOT_FOUND", http.StatusNotFound)
		return
	}

	response.NewSuccessResponse(w, data, http.StatusOK)
}
