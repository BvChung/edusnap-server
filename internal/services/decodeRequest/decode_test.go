package decoderequest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/ocr/internal/models"
)

func TestDecodeReqBody(t *testing.T) {
	t.Run("Test decoding", func(t *testing.T) {
		userUUID, _ := uuid.Parse("d7e2d1e3-30ed-4332-b26c-45314f33bad3")
		message := "Test message" 
		encoded_images := []string{"data:image/png;base64,SGVsbG8sIOS4lueVjA=="}
		newMessageReq := models.Message{UserID: &userUUID, Message: message, EncodedImages: encoded_images} 
		data, _ := json.Marshal(newMessageReq)

		r := httptest.NewRequest(http.MethodPost, "/api/message", bytes.NewBuffer(data))

		var decodedBody models.Message
		if err := DecodeRequestBody(r, &decodedBody); err != nil {
			t.Fatalf(err.Error())
		}

		t.Log(decodedBody)
	})
}