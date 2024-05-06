package message

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/ocr/cmd/models"
	"github.com/ocr/cmd/services/vertexai"
)

func ProcessMessageData(messageRequestBody *models.Message) ([]models.ReturnedMessage, error) {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	errChan := make(chan error, len(messageRequestBody.EncodedImages))

	var data []models.ReturnedMessage = make([]models.ReturnedMessage, 5)
	var images []*models.Image = make([]*models.Image, 0, 10)

	for _, encImg := range messageRequestBody.EncodedImages {
		wg.Add(1)

		go decodeImage(&encImg, &images, &mu, &wg, errChan)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return data, fmt.Errorf("error processing uploaded images")
		}
	}

	if len(images) > 0 {
		for i := 0; i < len(images); i++ {
			id := uuid.New()

			returnedMessage := models.ReturnedMessage{ID: &id, Message: messageRequestBody.Message, DecodedImage: string((*images[i].RawEncoding))}
			data = append(data, returnedMessage)
		}
	}

	return data, nil
}

func decodeImage(encodedImg *string, images *[]*models.Image, mu *sync.Mutex, wg *sync.WaitGroup, errCh chan<- error) {
	defer func() {
		mu.Unlock()
		wg.Done()
	}()

	mu.Lock()
	b64, mimeType, err := vertexai.ExtractBase64(*encodedImg)

	if err != nil {
		errCh <- err
		return
	}

	processedImg := &models.Image{MimeType: mimeType, Base64: &b64, RawEncoding: nil}

	if err = processedImg.DecodeBase64(); err != nil {
		errCh <- err
		return
	}

	*images = append(*images, processedImg)
}
