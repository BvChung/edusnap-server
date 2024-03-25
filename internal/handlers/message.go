package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"cloud.google.com/go/vertexai/genai"
	"github.com/google/uuid"
	"github.com/ocr/internal/response"
	"github.com/ocr/internal/vertexai"
	"github.com/supabase-community/supabase-go"
)

type Message struct {
	UserID           *uuid.UUID `json:"user_id"`
	Message      string     `json:"message"`
	EncodedImages []string     `json:"encoded_images"`
}

type Image struct {
	MimeType string
	Base64 *string
	RawEncoding *[]byte
}

func (img *Image) DecodeBase64() error {
	if img == nil {
		return fmt.Errorf("img nil pointer exception")
	}

	rawEncoding, err := base64.StdEncoding.DecodeString(*img.Base64)

	if err != nil {
		return fmt.Errorf("unable to decode base64 string: %w", err)
	}

	img.RawEncoding = &rawEncoding

	return nil
}

type ReturnedMessage struct {
	ID           *uuid.UUID `json:"id"`
	Message      string     `json:"message"`
	DecodedImage string     `json:"decoded_image"`
}

type MessagesHandler struct {
	DBClient *supabase.Client
	VertexAIClient *genai.Client
}

func (mh *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		GetMessages(mh.DBClient, w, r)
	case "POST":
		CreateMessage(mh.DBClient, w, r)
	}
}

func NewMessagesHandler(s *supabase.Client, v *genai.Client) *MessagesHandler {
	return &MessagesHandler{DBClient: s, VertexAIClient: v}
}

func GetMessages(client *supabase.Client, w http.ResponseWriter, r *http.Request) {
	var data []Message

	if _, err := client.From("messages").Select("*", "", false).ExecuteTo(&data); err != nil {
		response.NewErrorResponse(w, "unable to get messages", response.ServerError, http.StatusInternalServerError)
		return
	}

	response.NewSuccessResponse(w, data, http.StatusOK)
}

func CreateMessage(s *supabase.Client, w http.ResponseWriter, r *http.Request) {
	var newMessage Message

	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		log.Println("Failed to decode request body to JSON: ", err.Error())
		response.NewErrorResponse(w, "Malformed request body", "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	errCh := make(chan error, len(newMessage.EncodedImages))
	
	data := []ReturnedMessage{}
	var images []*Image = make([]*Image, 0, 10)

	for _, encImg := range newMessage.EncodedImages{
		wg.Add(1)

		go processImage(&encImg, &images, &mu, &wg, errCh)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh{
		if err != nil {
			log.Println(err.Error())
			response.NewErrorResponse(w, "Error processing uploaded images", response.BadRequest, http.StatusBadRequest)
			return
		}
	}

	if len(images) > 0 {
		for i := 0; i < len(images); i++ {
			id := uuid.New()

			returnedMessage := ReturnedMessage{ID: &id, Message: newMessage.Message, DecodedImage: string((*images[i].RawEncoding))}
			data = append(data, returnedMessage)
		}
	}

	response.NewSuccessResponse(w, data, http.StatusOK)
}

func processImage(encodedImg *string, images *[]*Image, mu *sync.Mutex, wg *sync.WaitGroup, errCh chan<- error){
	defer func(){
		mu.Unlock()
		wg.Done()
	}()

	mu.Lock()
	b64, mimeType, err := vertexai.ExtractBase64(*encodedImg)

	if err != nil {
		errCh <- err
		return
	}

	processedImg := &Image{MimeType: mimeType, Base64: &b64, RawEncoding: nil}

	if err = processedImg.DecodeBase64(); err != nil {
		errCh <- err
		return
	}

	*images = append(*images, processedImg)
}