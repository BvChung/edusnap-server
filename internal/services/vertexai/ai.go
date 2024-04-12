package vertexai

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

func CreateVertexClient() (*genai.Client, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	location := os.Getenv("GOOGLE_CLOUD_REGION")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, projectID, location)

	if err != nil {
		return client, fmt.Errorf("unable to create vertex ai client: %w", err)
	}

	return client, nil
}

func GenerateAIServiceRequest(client *genai.Client) {

}

func MakeGeminiRequest(format string, message string, decodedImage []byte) ([]genai.Part, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	location := os.Getenv("GOOGLE_CLOUD_REGION")
	modelName := os.Getenv("GOOGLE_CLOUD_VERTEX_MODEL_NAME")
	var temperature float32 = 0.0

	var img genai.Blob = genai.ImageData(format, decodedImage)

	prompt := []genai.Part{
		img,
		genai.Text(message),
	}

	content, err := generateMultimodalContent(prompt, projectID, location, modelName, float32(temperature))
	if err != nil {
		return content, err
	}

	return content, nil
}

// generateMultimodalContent provide a generated response using multimodal input
func generateMultimodalContent(parts []genai.Part, projectID, location, modelName string, temperature float32) ([]genai.Part, error) {
	var content []genai.Part
	ctx := context.Background()

	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return content, fmt.Errorf("unable to create vertex ai client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	model.SetTemperature(temperature)

	res, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		return content, fmt.Errorf("unable to generate contents: %w", err)
	}

	content = getResponse(res)

	return content, nil
}

func getResponse(res *genai.GenerateContentResponse) []genai.Part {
	geminiRes := []genai.Part{}

	for _, cand := range res.Candidates {
		if cand.Content.Role == "model" {
			geminiRes = append(geminiRes, cand.Content.Parts...)
		}
	}

	return geminiRes
}
