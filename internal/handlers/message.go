package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ocr/internal/format"
	"github.com/ocr/internal/vertexai"
)

type Message struct {
	ID           *uuid.UUID `json:"id"`
	Message      string     `json:"message"`
	EncodedImage string     `json:"encoded_image"`
}

type ReturnedMessage struct {
	ID           *uuid.UUID `json:"id"`
	Message      string     `json:"message"`
	DecodedImage string     `json:"decoded_image"`
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	var newMessage Message

	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		format.NewErrorResponse(w, "Internal server error", "SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	b64, mimeType, err := vertexai.ExtractBase64(newMessage.EncodedImage)
	if err != nil {
		format.NewErrorResponse(w, "Invalid base64 encoded string", "INVALID_BASE64", http.StatusBadRequest)
		return
	}

	fmt.Println(mimeType)

	data := []ReturnedMessage{}

	if newMessage.EncodedImage != "" {
		decodedImg, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			fmt.Println("decode error:", err)
			format.NewErrorResponse(w, "Unable to decode base64 encoded image", "DECODE_ERROR", http.StatusInternalServerError)
			return
		}

		returnedMessage := ReturnedMessage{ID: newMessage.ID, Message: newMessage.Message, DecodedImage: string(decodedImg)}
		data = append(data, returnedMessage)
	}

	format.NewSuccessResponse(w, data, http.StatusOK)
}
