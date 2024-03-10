package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/internal/format"
)

type Message struct{
	ID *uuid.UUID `json:"id"`
	Message string `json:"message"`
	EncodedImage string `json:"encoded_image"`
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	var newMessage Message

	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		format.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	data := []Message{}

	format.NewSuccessResponse(w, data, http.StatusOK)
}