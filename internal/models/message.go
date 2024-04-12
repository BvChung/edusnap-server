package models

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

type Message struct {
	UserID        *uuid.UUID `json:"user_id"`
	Message       string     `json:"message"`
	EncodedImages []string   `json:"encoded_images"`
}

type Image struct {
	MimeType    string
	Base64      *string
	RawEncoding *[]byte
}

func (img *Image) DecodeBase64() error {
	if img == nil {
		return fmt.Errorf("Image struct is nil")
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