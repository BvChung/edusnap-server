package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	"github.com/ocr/cmd/models"
	"github.com/ocr/cmd/response"
	"github.com/ocr/cmd/services/message"
	"github.com/supabase-community/supabase-go"
)

type MessagesHandler struct {
	DBClient       *supabase.Client
	VertexAIClient *genai.Client
}

func (mh *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		getMessages(mh.DBClient, w, r)
	case "POST":
		handleMessageCreation(mh.DBClient, w, r)
	}
}

func NewMessagesHandler(s *supabase.Client, v *genai.Client) *MessagesHandler {
	return &MessagesHandler{DBClient: s, VertexAIClient: v}
}

func getMessages(client *supabase.Client, w http.ResponseWriter, r *http.Request) {
	var data []models.Message

	if _, err := client.From("messages").Select("*", "", false).ExecuteTo(&data); err != nil {
		response.NewErrorResponse(w, "unable to get messages", response.ServerError, http.StatusInternalServerError)
		return
	}

	response.NewSuccessResponse(w, data, http.StatusOK)
}

func handleMessageCreation(s *supabase.Client, w http.ResponseWriter, r *http.Request) {
	var messageRequestBody models.Message

	if err := json.NewDecoder(r.Body).Decode(&messageRequestBody); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		response.NewErrorResponse(w, "Malformed request body", "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	data, err := message.ProcessMessageData(&messageRequestBody)
	if err != nil {
		log.Println(err.Error())
		response.NewErrorResponse(w, "Error processing uploaded images", response.BadRequest, http.StatusBadRequest)
		return
	}

	response.NewSuccessResponse(w, data, http.StatusOK)
}
